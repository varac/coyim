package definitions

func init() {
	add(`AccountRegistration`, &defAccountRegistration{})
}

type defAccountRegistration struct{}

func (*defAccountRegistration) String() string {
	return `<interface>
  <object class="GtkDialog" id="dialog">
    <property name="title" translatable="yes">Choose a server to register your account</property>
    <signal name="close" handler="close-handler" />

    <child internal-child="vbox">
      <object class="GtkBox" id="Vbox">
        <property name="margin">10</property>
        <child>
          <object class="GtkGrid" id="grid">
            <property name="margin-bottom">10</property>
            <property name="row-spacing">12</property>
            <property name="column-spacing">6</property>

            <child>
              <object class="GtkLabel" id="server-label">
                <property name="label" translatable="yes">Server</property>
                <property name="justify">GTK_JUSTIFY_RIGHT</property>
              </object>
              <packing>
                <property name="left-attach">0</property>
                <property name="top-attach">0</property>
              </packing>
            </child>
            <child>
              <object class="GtkComboBoxText" id="server">
                <property name="has-entry">True</property>
              </object>
              <packing>
                <property name="left-attach">1</property>
                <property name="top-attach">0</property>
              </packing>
            </child>
          </object>
        </child>

        <child internal-child="action_area">
          <object class="GtkButtonBox" id="bbox">
            <property name="orientation">GTK_ORIENTATION_HORIZONTAL</property>
            <child>
              <object class="GtkButton" id="btn-cancel">
                <property name="label" translatable="yes">Cancel</property>
              </object>
            </child>
            <child>
              <object class="GtkButton" id="btn-apply">
                <property name="label" translatable="yes">Apply</property>
                <property name="can-default">true</property>
              </object>
            </child>
          </object>
        </child>
      </object>
    </child>

    <action-widgets>
      <action-widget response="cancel">btn-cancel</action-widget>
      <action-widget response="apply" default="true">btn-apply</action-widget>
    </action-widgets>
  </object>
</interface>
`
}
