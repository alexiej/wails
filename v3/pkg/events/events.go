package events

type ApplicationEventType uint
type WindowEventType uint

var Common = newCommonEvents()

type commonEvents struct {
	ApplicationStarted ApplicationEventType
	WindowMaximise     WindowEventType
	WindowUnMaximise   WindowEventType
	WindowFullscreen   WindowEventType
	WindowUnFullscreen WindowEventType
	WindowRestore      WindowEventType
	WindowMinimise     WindowEventType
	WindowUnMinimise   WindowEventType
	WindowClosing      WindowEventType
	WindowZoom         WindowEventType
	WindowZoomIn       WindowEventType
	WindowZoomOut      WindowEventType
	WindowZoomReset    WindowEventType
	WindowFocus        WindowEventType
	WindowLostFocus    WindowEventType
	WindowShow         WindowEventType
	WindowHide         WindowEventType
	WindowDPIChanged   WindowEventType
	WindowFilesDropped WindowEventType
	WindowRuntimeReady WindowEventType
	ThemeChanged       ApplicationEventType
	WindowDidMove      WindowEventType
	WindowDidResize    WindowEventType
	WindowDragDrop     WindowEventType
	WindowDragEnter    WindowEventType
	WindowDragLeave    WindowEventType
	WindowDragOver     WindowEventType
}

func newCommonEvents() commonEvents {
	return commonEvents{
		ApplicationStarted: 1187,
		WindowMaximise:     1188,
		WindowUnMaximise:   1189,
		WindowFullscreen:   1190,
		WindowUnFullscreen: 1191,
		WindowRestore:      1192,
		WindowMinimise:     1193,
		WindowUnMinimise:   1194,
		WindowClosing:      1195,
		WindowZoom:         1196,
		WindowZoomIn:       1197,
		WindowZoomOut:      1198,
		WindowZoomReset:    1199,
		WindowFocus:        1200,
		WindowLostFocus:    1201,
		WindowShow:         1202,
		WindowHide:         1203,
		WindowDPIChanged:   1204,
		WindowFilesDropped: 1205,
		WindowRuntimeReady: 1206,
		ThemeChanged:       1207,
		WindowDidMove:      1208,
		WindowDidResize:    1209,
		WindowDragDrop:     1210,
		WindowDragEnter:    1211,
		WindowDragLeave:    1212,
		WindowDragOver:     1213,
	}
}

var Linux = newLinuxEvents()

type linuxEvents struct {
	SystemThemeChanged ApplicationEventType
	WindowLoadChanged  WindowEventType
	WindowDeleteEvent  WindowEventType
	WindowDidMove      WindowEventType
	WindowDidResize    WindowEventType
	WindowFocusIn      WindowEventType
	WindowFocusOut     WindowEventType
	WindowDragDrop     WindowEventType
	WindowDragBegin    WindowEventType
	WindowDragEnd      WindowEventType
	WindowDragLeave    WindowEventType
	ApplicationStartup ApplicationEventType
}

func newLinuxEvents() linuxEvents {
	return linuxEvents{
		SystemThemeChanged: 1024,
		WindowLoadChanged:  1025,
		WindowDeleteEvent:  1026,
		WindowDidMove:      1027,
		WindowDidResize:    1028,
		WindowFocusIn:      1029,
		WindowFocusOut:     1030,
		WindowDragDrop:     1031,
		WindowDragBegin:    1032,
		WindowDragEnd:      1033,
		WindowDragLeave:    1034,
		ApplicationStartup: 1035,
	}
}

var Mac = newMacEvents()

type macEvents struct {
	ApplicationDidBecomeActive                              ApplicationEventType
	ApplicationDidChangeBackingProperties                   ApplicationEventType
	ApplicationDidChangeEffectiveAppearance                 ApplicationEventType
	ApplicationDidChangeIcon                                ApplicationEventType
	ApplicationDidChangeOcclusionState                      ApplicationEventType
	ApplicationDidChangeScreenParameters                    ApplicationEventType
	ApplicationDidChangeStatusBarFrame                      ApplicationEventType
	ApplicationDidChangeStatusBarOrientation                ApplicationEventType
	ApplicationDidFinishLaunching                           ApplicationEventType
	ApplicationDidHide                                      ApplicationEventType
	ApplicationDidResignActiveNotification                  ApplicationEventType
	ApplicationDidUnhide                                    ApplicationEventType
	ApplicationDidUpdate                                    ApplicationEventType
	ApplicationWillBecomeActive                             ApplicationEventType
	ApplicationWillFinishLaunching                          ApplicationEventType
	ApplicationWillHide                                     ApplicationEventType
	ApplicationWillResignActive                             ApplicationEventType
	ApplicationWillTerminate                                ApplicationEventType
	ApplicationWillUnhide                                   ApplicationEventType
	ApplicationWillUpdate                                   ApplicationEventType
	ApplicationDidChangeTheme                               ApplicationEventType
	ApplicationShouldHandleReopen                           ApplicationEventType
	WindowDidBecomeKey                                      WindowEventType
	WindowDidBecomeMain                                     WindowEventType
	WindowDidBeginSheet                                     WindowEventType
	WindowDidChangeAlpha                                    WindowEventType
	WindowDidChangeBackingLocation                          WindowEventType
	WindowDidChangeBackingProperties                        WindowEventType
	WindowDidChangeCollectionBehavior                       WindowEventType
	WindowDidChangeEffectiveAppearance                      WindowEventType
	WindowDidChangeOcclusionState                           WindowEventType
	WindowDidChangeOrderingMode                             WindowEventType
	WindowDidChangeScreen                                   WindowEventType
	WindowDidChangeScreenParameters                         WindowEventType
	WindowDidChangeScreenProfile                            WindowEventType
	WindowDidChangeScreenSpace                              WindowEventType
	WindowDidChangeScreenSpaceProperties                    WindowEventType
	WindowDidChangeSharingType                              WindowEventType
	WindowDidChangeSpace                                    WindowEventType
	WindowDidChangeSpaceOrderingMode                        WindowEventType
	WindowDidChangeTitle                                    WindowEventType
	WindowDidChangeToolbar                                  WindowEventType
	WindowDidChangeVisibility                               WindowEventType
	WindowDidDeminiaturize                                  WindowEventType
	WindowDidEndSheet                                       WindowEventType
	WindowDidEnterFullScreen                                WindowEventType
	WindowDidEnterVersionBrowser                            WindowEventType
	WindowDidExitFullScreen                                 WindowEventType
	WindowDidExitVersionBrowser                             WindowEventType
	WindowDidExpose                                         WindowEventType
	WindowDidFocus                                          WindowEventType
	WindowDidMiniaturize                                    WindowEventType
	WindowDidMove                                           WindowEventType
	WindowDidOrderOffScreen                                 WindowEventType
	WindowDidOrderOnScreen                                  WindowEventType
	WindowDidResignKey                                      WindowEventType
	WindowDidResignMain                                     WindowEventType
	WindowDidResize                                         WindowEventType
	WindowDidUpdate                                         WindowEventType
	WindowDidUpdateAlpha                                    WindowEventType
	WindowDidUpdateCollectionBehavior                       WindowEventType
	WindowDidUpdateCollectionProperties                     WindowEventType
	WindowDidUpdateShadow                                   WindowEventType
	WindowDidUpdateTitle                                    WindowEventType
	WindowDidUpdateToolbar                                  WindowEventType
	WindowDidUpdateVisibility                               WindowEventType
	WindowShouldClose                                       WindowEventType
	WindowWillBecomeKey                                     WindowEventType
	WindowWillBecomeMain                                    WindowEventType
	WindowWillBeginSheet                                    WindowEventType
	WindowWillChangeOrderingMode                            WindowEventType
	WindowWillClose                                         WindowEventType
	WindowWillDeminiaturize                                 WindowEventType
	WindowWillEnterFullScreen                               WindowEventType
	WindowWillEnterVersionBrowser                           WindowEventType
	WindowWillExitFullScreen                                WindowEventType
	WindowWillExitVersionBrowser                            WindowEventType
	WindowWillFocus                                         WindowEventType
	WindowWillMiniaturize                                   WindowEventType
	WindowWillMove                                          WindowEventType
	WindowWillOrderOffScreen                                WindowEventType
	WindowWillOrderOnScreen                                 WindowEventType
	WindowWillResignMain                                    WindowEventType
	WindowWillResize                                        WindowEventType
	WindowWillUnfocus                                       WindowEventType
	WindowWillUpdate                                        WindowEventType
	WindowWillUpdateAlpha                                   WindowEventType
	WindowWillUpdateCollectionBehavior                      WindowEventType
	WindowWillUpdateCollectionProperties                    WindowEventType
	WindowWillUpdateShadow                                  WindowEventType
	WindowWillUpdateTitle                                   WindowEventType
	WindowWillUpdateToolbar                                 WindowEventType
	WindowWillUpdateVisibility                              WindowEventType
	WindowWillUseStandardFrame                              WindowEventType
	MenuWillOpen                                            ApplicationEventType
	MenuDidOpen                                             ApplicationEventType
	MenuDidClose                                            ApplicationEventType
	MenuWillSendAction                                      ApplicationEventType
	MenuDidSendAction                                       ApplicationEventType
	MenuWillHighlightItem                                   ApplicationEventType
	MenuDidHighlightItem                                    ApplicationEventType
	MenuWillDisplayItem                                     ApplicationEventType
	MenuDidDisplayItem                                      ApplicationEventType
	MenuWillAddItem                                         ApplicationEventType
	MenuDidAddItem                                          ApplicationEventType
	MenuWillRemoveItem                                      ApplicationEventType
	MenuDidRemoveItem                                       ApplicationEventType
	MenuWillBeginTracking                                   ApplicationEventType
	MenuDidBeginTracking                                    ApplicationEventType
	MenuWillEndTracking                                     ApplicationEventType
	MenuDidEndTracking                                      ApplicationEventType
	MenuWillUpdate                                          ApplicationEventType
	MenuDidUpdate                                           ApplicationEventType
	MenuWillPopUp                                           ApplicationEventType
	MenuDidPopUp                                            ApplicationEventType
	MenuWillSendActionToItem                                ApplicationEventType
	MenuDidSendActionToItem                                 ApplicationEventType
	WebViewDidStartProvisionalNavigation                    WindowEventType
	WebViewDidReceiveServerRedirectForProvisionalNavigation WindowEventType
	WebViewDidFinishNavigation                              WindowEventType
	WebViewDidCommitNavigation                              WindowEventType
	WindowFileDraggingEntered                               WindowEventType
	WindowFileDraggingPerformed                             WindowEventType
	WindowFileDraggingExited                                WindowEventType
}

func newMacEvents() macEvents {
	return macEvents{
		ApplicationDidBecomeActive:               1036,
		ApplicationDidChangeBackingProperties:    1037,
		ApplicationDidChangeEffectiveAppearance:  1038,
		ApplicationDidChangeIcon:                 1039,
		ApplicationDidChangeOcclusionState:       1040,
		ApplicationDidChangeScreenParameters:     1041,
		ApplicationDidChangeStatusBarFrame:       1042,
		ApplicationDidChangeStatusBarOrientation: 1043,
		ApplicationDidFinishLaunching:            1044,
		ApplicationDidHide:                       1045,
		ApplicationDidResignActiveNotification:   1046,
		ApplicationDidUnhide:                     1047,
		ApplicationDidUpdate:                     1048,
		ApplicationWillBecomeActive:              1049,
		ApplicationWillFinishLaunching:           1050,
		ApplicationWillHide:                      1051,
		ApplicationWillResignActive:              1052,
		ApplicationWillTerminate:                 1053,
		ApplicationWillUnhide:                    1054,
		ApplicationWillUpdate:                    1055,
		ApplicationDidChangeTheme:                1056,
		ApplicationShouldHandleReopen:            1057,
		WindowDidBecomeKey:                       1058,
		WindowDidBecomeMain:                      1059,
		WindowDidBeginSheet:                      1060,
		WindowDidChangeAlpha:                     1061,
		WindowDidChangeBackingLocation:           1062,
		WindowDidChangeBackingProperties:         1063,
		WindowDidChangeCollectionBehavior:        1064,
		WindowDidChangeEffectiveAppearance:       1065,
		WindowDidChangeOcclusionState:            1066,
		WindowDidChangeOrderingMode:              1067,
		WindowDidChangeScreen:                    1068,
		WindowDidChangeScreenParameters:          1069,
		WindowDidChangeScreenProfile:             1070,
		WindowDidChangeScreenSpace:               1071,
		WindowDidChangeScreenSpaceProperties:     1072,
		WindowDidChangeSharingType:               1073,
		WindowDidChangeSpace:                     1074,
		WindowDidChangeSpaceOrderingMode:         1075,
		WindowDidChangeTitle:                     1076,
		WindowDidChangeToolbar:                   1077,
		WindowDidChangeVisibility:                1078,
		WindowDidDeminiaturize:                   1079,
		WindowDidEndSheet:                        1080,
		WindowDidEnterFullScreen:                 1081,
		WindowDidEnterVersionBrowser:             1082,
		WindowDidExitFullScreen:                  1083,
		WindowDidExitVersionBrowser:              1084,
		WindowDidExpose:                          1085,
		WindowDidFocus:                           1086,
		WindowDidMiniaturize:                     1087,
		WindowDidMove:                            1088,
		WindowDidOrderOffScreen:                  1089,
		WindowDidOrderOnScreen:                   1090,
		WindowDidResignKey:                       1091,
		WindowDidResignMain:                      1092,
		WindowDidResize:                          1093,
		WindowDidUpdate:                          1094,
		WindowDidUpdateAlpha:                     1095,
		WindowDidUpdateCollectionBehavior:        1096,
		WindowDidUpdateCollectionProperties:      1097,
		WindowDidUpdateShadow:                    1098,
		WindowDidUpdateTitle:                     1099,
		WindowDidUpdateToolbar:                   1100,
		WindowDidUpdateVisibility:                1101,
		WindowShouldClose:                        1102,
		WindowWillBecomeKey:                      1103,
		WindowWillBecomeMain:                     1104,
		WindowWillBeginSheet:                     1105,
		WindowWillChangeOrderingMode:             1106,
		WindowWillClose:                          1107,
		WindowWillDeminiaturize:                  1108,
		WindowWillEnterFullScreen:                1109,
		WindowWillEnterVersionBrowser:            1110,
		WindowWillExitFullScreen:                 1111,
		WindowWillExitVersionBrowser:             1112,
		WindowWillFocus:                          1113,
		WindowWillMiniaturize:                    1114,
		WindowWillMove:                           1115,
		WindowWillOrderOffScreen:                 1116,
		WindowWillOrderOnScreen:                  1117,
		WindowWillResignMain:                     1118,
		WindowWillResize:                         1119,
		WindowWillUnfocus:                        1120,
		WindowWillUpdate:                         1121,
		WindowWillUpdateAlpha:                    1122,
		WindowWillUpdateCollectionBehavior:       1123,
		WindowWillUpdateCollectionProperties:     1124,
		WindowWillUpdateShadow:                   1125,
		WindowWillUpdateTitle:                    1126,
		WindowWillUpdateToolbar:                  1127,
		WindowWillUpdateVisibility:               1128,
		WindowWillUseStandardFrame:               1129,
		MenuWillOpen:                             1130,
		MenuDidOpen:                              1131,
		MenuDidClose:                             1132,
		MenuWillSendAction:                       1133,
		MenuDidSendAction:                        1134,
		MenuWillHighlightItem:                    1135,
		MenuDidHighlightItem:                     1136,
		MenuWillDisplayItem:                      1137,
		MenuDidDisplayItem:                       1138,
		MenuWillAddItem:                          1139,
		MenuDidAddItem:                           1140,
		MenuWillRemoveItem:                       1141,
		MenuDidRemoveItem:                        1142,
		MenuWillBeginTracking:                    1143,
		MenuDidBeginTracking:                     1144,
		MenuWillEndTracking:                      1145,
		MenuDidEndTracking:                       1146,
		MenuWillUpdate:                           1147,
		MenuDidUpdate:                            1148,
		MenuWillPopUp:                            1149,
		MenuDidPopUp:                             1150,
		MenuWillSendActionToItem:                 1151,
		MenuDidSendActionToItem:                  1152,
		WebViewDidStartProvisionalNavigation:     1153,
		WebViewDidReceiveServerRedirectForProvisionalNavigation: 1154,
		WebViewDidFinishNavigation:                              1155,
		WebViewDidCommitNavigation:                              1156,
		WindowFileDraggingEntered:                               1157,
		WindowFileDraggingPerformed:                             1158,
		WindowFileDraggingExited:                                1159,
	}
}

var Windows = newWindowsEvents()

type windowsEvents struct {
	SystemThemeChanged         ApplicationEventType
	APMPowerStatusChange       ApplicationEventType
	APMSuspend                 ApplicationEventType
	APMResumeAutomatic         ApplicationEventType
	APMResumeSuspend           ApplicationEventType
	APMPowerSettingChange      ApplicationEventType
	ApplicationStarted         ApplicationEventType
	WebViewNavigationCompleted WindowEventType
	WindowInactive             WindowEventType
	WindowActive               WindowEventType
	WindowClickActive          WindowEventType
	WindowMaximise             WindowEventType
	WindowUnMaximise           WindowEventType
	WindowFullscreen           WindowEventType
	WindowUnFullscreen         WindowEventType
	WindowRestore              WindowEventType
	WindowMinimise             WindowEventType
	WindowUnMinimise           WindowEventType
	WindowClose                WindowEventType
	WindowSetFocus             WindowEventType
	WindowKillFocus            WindowEventType
	WindowDragDrop             WindowEventType
	WindowDragEnter            WindowEventType
	WindowDragLeave            WindowEventType
	WindowDragOver             WindowEventType
	WindowDidMove              WindowEventType
	WindowDidResize            WindowEventType
}

func newWindowsEvents() windowsEvents {
	return windowsEvents{
		SystemThemeChanged:         1160,
		APMPowerStatusChange:       1161,
		APMSuspend:                 1162,
		APMResumeAutomatic:         1163,
		APMResumeSuspend:           1164,
		APMPowerSettingChange:      1165,
		ApplicationStarted:         1166,
		WebViewNavigationCompleted: 1167,
		WindowInactive:             1168,
		WindowActive:               1169,
		WindowClickActive:          1170,
		WindowMaximise:             1171,
		WindowUnMaximise:           1172,
		WindowFullscreen:           1173,
		WindowUnFullscreen:         1174,
		WindowRestore:              1175,
		WindowMinimise:             1176,
		WindowUnMinimise:           1177,
		WindowClose:                1178,
		WindowSetFocus:             1179,
		WindowKillFocus:            1180,
		WindowDragDrop:             1181,
		WindowDragEnter:            1182,
		WindowDragLeave:            1183,
		WindowDragOver:             1184,
		WindowDidMove:              1185,
		WindowDidResize:            1186,
	}
}

func JSEvent(event uint) string {
	return eventToJS[event]
}

var eventToJS = map[uint]string{
	1024: "linux:SystemThemeChanged",
	1025: "linux:WindowLoadChanged",
	1026: "linux:WindowDeleteEvent",
	1027: "linux:WindowDidMove",
	1028: "linux:WindowDidResize",
	1029: "linux:WindowFocusIn",
	1030: "linux:WindowFocusOut",
	1031: "linux:WindowDragDrop",
	1032: "linux:WindowDragBegin",
	1033: "linux:WindowDragEnd",
	1034: "linux:WindowDragLeave",
	1035: "linux:ApplicationStartup",
	1036: "mac:ApplicationDidBecomeActive",
	1037: "mac:ApplicationDidChangeBackingProperties",
	1038: "mac:ApplicationDidChangeEffectiveAppearance",
	1039: "mac:ApplicationDidChangeIcon",
	1040: "mac:ApplicationDidChangeOcclusionState",
	1041: "mac:ApplicationDidChangeScreenParameters",
	1042: "mac:ApplicationDidChangeStatusBarFrame",
	1043: "mac:ApplicationDidChangeStatusBarOrientation",
	1044: "mac:ApplicationDidFinishLaunching",
	1045: "mac:ApplicationDidHide",
	1046: "mac:ApplicationDidResignActiveNotification",
	1047: "mac:ApplicationDidUnhide",
	1048: "mac:ApplicationDidUpdate",
	1049: "mac:ApplicationWillBecomeActive",
	1050: "mac:ApplicationWillFinishLaunching",
	1051: "mac:ApplicationWillHide",
	1052: "mac:ApplicationWillResignActive",
	1053: "mac:ApplicationWillTerminate",
	1054: "mac:ApplicationWillUnhide",
	1055: "mac:ApplicationWillUpdate",
	1056: "mac:ApplicationDidChangeTheme!",
	1057: "mac:ApplicationShouldHandleReopen!",
	1058: "mac:WindowDidBecomeKey",
	1059: "mac:WindowDidBecomeMain",
	1060: "mac:WindowDidBeginSheet",
	1061: "mac:WindowDidChangeAlpha",
	1062: "mac:WindowDidChangeBackingLocation",
	1063: "mac:WindowDidChangeBackingProperties",
	1064: "mac:WindowDidChangeCollectionBehavior",
	1065: "mac:WindowDidChangeEffectiveAppearance",
	1066: "mac:WindowDidChangeOcclusionState",
	1067: "mac:WindowDidChangeOrderingMode",
	1068: "mac:WindowDidChangeScreen",
	1069: "mac:WindowDidChangeScreenParameters",
	1070: "mac:WindowDidChangeScreenProfile",
	1071: "mac:WindowDidChangeScreenSpace",
	1072: "mac:WindowDidChangeScreenSpaceProperties",
	1073: "mac:WindowDidChangeSharingType",
	1074: "mac:WindowDidChangeSpace",
	1075: "mac:WindowDidChangeSpaceOrderingMode",
	1076: "mac:WindowDidChangeTitle",
	1077: "mac:WindowDidChangeToolbar",
	1078: "mac:WindowDidChangeVisibility",
	1079: "mac:WindowDidDeminiaturize",
	1080: "mac:WindowDidEndSheet",
	1081: "mac:WindowDidEnterFullScreen",
	1082: "mac:WindowDidEnterVersionBrowser",
	1083: "mac:WindowDidExitFullScreen",
	1084: "mac:WindowDidExitVersionBrowser",
	1085: "mac:WindowDidExpose",
	1086: "mac:WindowDidFocus",
	1087: "mac:WindowDidMiniaturize",
	1088: "mac:WindowDidMove",
	1089: "mac:WindowDidOrderOffScreen",
	1090: "mac:WindowDidOrderOnScreen",
	1091: "mac:WindowDidResignKey",
	1092: "mac:WindowDidResignMain",
	1093: "mac:WindowDidResize",
	1094: "mac:WindowDidUpdate",
	1095: "mac:WindowDidUpdateAlpha",
	1096: "mac:WindowDidUpdateCollectionBehavior",
	1097: "mac:WindowDidUpdateCollectionProperties",
	1098: "mac:WindowDidUpdateShadow",
	1099: "mac:WindowDidUpdateTitle",
	1100: "mac:WindowDidUpdateToolbar",
	1101: "mac:WindowDidUpdateVisibility",
	1102: "mac:WindowShouldClose!",
	1103: "mac:WindowWillBecomeKey",
	1104: "mac:WindowWillBecomeMain",
	1105: "mac:WindowWillBeginSheet",
	1106: "mac:WindowWillChangeOrderingMode",
	1107: "mac:WindowWillClose",
	1108: "mac:WindowWillDeminiaturize",
	1109: "mac:WindowWillEnterFullScreen",
	1110: "mac:WindowWillEnterVersionBrowser",
	1111: "mac:WindowWillExitFullScreen",
	1112: "mac:WindowWillExitVersionBrowser",
	1113: "mac:WindowWillFocus",
	1114: "mac:WindowWillMiniaturize",
	1115: "mac:WindowWillMove",
	1116: "mac:WindowWillOrderOffScreen",
	1117: "mac:WindowWillOrderOnScreen",
	1118: "mac:WindowWillResignMain",
	1119: "mac:WindowWillResize",
	1120: "mac:WindowWillUnfocus",
	1121: "mac:WindowWillUpdate",
	1122: "mac:WindowWillUpdateAlpha",
	1123: "mac:WindowWillUpdateCollectionBehavior",
	1124: "mac:WindowWillUpdateCollectionProperties",
	1125: "mac:WindowWillUpdateShadow",
	1126: "mac:WindowWillUpdateTitle",
	1127: "mac:WindowWillUpdateToolbar",
	1128: "mac:WindowWillUpdateVisibility",
	1129: "mac:WindowWillUseStandardFrame",
	1130: "mac:MenuWillOpen",
	1131: "mac:MenuDidOpen",
	1132: "mac:MenuDidClose",
	1133: "mac:MenuWillSendAction",
	1134: "mac:MenuDidSendAction",
	1135: "mac:MenuWillHighlightItem",
	1136: "mac:MenuDidHighlightItem",
	1137: "mac:MenuWillDisplayItem",
	1138: "mac:MenuDidDisplayItem",
	1139: "mac:MenuWillAddItem",
	1140: "mac:MenuDidAddItem",
	1141: "mac:MenuWillRemoveItem",
	1142: "mac:MenuDidRemoveItem",
	1143: "mac:MenuWillBeginTracking",
	1144: "mac:MenuDidBeginTracking",
	1145: "mac:MenuWillEndTracking",
	1146: "mac:MenuDidEndTracking",
	1147: "mac:MenuWillUpdate",
	1148: "mac:MenuDidUpdate",
	1149: "mac:MenuWillPopUp",
	1150: "mac:MenuDidPopUp",
	1151: "mac:MenuWillSendActionToItem",
	1152: "mac:MenuDidSendActionToItem",
	1153: "mac:WebViewDidStartProvisionalNavigation",
	1154: "mac:WebViewDidReceiveServerRedirectForProvisionalNavigation",
	1155: "mac:WebViewDidFinishNavigation",
	1156: "mac:WebViewDidCommitNavigation",
	1157: "mac:WindowFileDraggingEntered",
	1158: "mac:WindowFileDraggingPerformed",
	1159: "mac:WindowFileDraggingExited",
	1160: "windows:SystemThemeChanged",
	1161: "windows:APMPowerStatusChange",
	1162: "windows:APMSuspend",
	1163: "windows:APMResumeAutomatic",
	1164: "windows:APMResumeSuspend",
	1165: "windows:APMPowerSettingChange",
	1166: "windows:ApplicationStarted",
	1167: "windows:WebViewNavigationCompleted",
	1168: "windows:WindowInactive",
	1169: "windows:WindowActive",
	1170: "windows:WindowClickActive",
	1171: "windows:WindowMaximise",
	1172: "windows:WindowUnMaximise",
	1173: "windows:WindowFullscreen",
	1174: "windows:WindowUnFullscreen",
	1175: "windows:WindowRestore",
	1176: "windows:WindowMinimise",
	1177: "windows:WindowUnMinimise",
	1178: "windows:WindowClose",
	1179: "windows:WindowSetFocus",
	1180: "windows:WindowKillFocus",
	1181: "windows:WindowDragDrop",
	1182: "windows:WindowDragEnter",
	1183: "windows:WindowDragLeave",
	1184: "windows:WindowDragOver",
	1185: "windows:WindowDidMove",
	1186: "windows:WindowDidResize",
	1187: "common:ApplicationStarted",
	1188: "common:WindowMaximise",
	1189: "common:WindowUnMaximise",
	1190: "common:WindowFullscreen",
	1191: "common:WindowUnFullscreen",
	1192: "common:WindowRestore",
	1193: "common:WindowMinimise",
	1194: "common:WindowUnMinimise",
	1195: "common:WindowClosing",
	1196: "common:WindowZoom",
	1197: "common:WindowZoomIn",
	1198: "common:WindowZoomOut",
	1199: "common:WindowZoomReset",
	1200: "common:WindowFocus",
	1201: "common:WindowLostFocus",
	1202: "common:WindowShow",
	1203: "common:WindowHide",
	1204: "common:WindowDPIChanged",
	1205: "common:WindowFilesDropped",
	1206: "common:WindowRuntimeReady",
	1207: "common:ThemeChanged",
	1208: "common:WindowDidMove",
	1209: "common:WindowDidResize",
	1210: "common:WindowDragDrop",
	1211: "common:WindowDragEnter",
	1212: "common:WindowDragLeave",
	1213: "common:WindowDragOver",
}
