// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetworkReachableAnalysisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliUid(v int64) *GetNetworkReachableAnalysisResponseBody
	GetAliUid() *int64
	SetCreateTime(v string) *GetNetworkReachableAnalysisResponseBody
	GetCreateTime() *string
	SetNetworkPathId(v string) *GetNetworkReachableAnalysisResponseBody
	GetNetworkPathId() *string
	SetNetworkPathParameter(v string) *GetNetworkReachableAnalysisResponseBody
	GetNetworkPathParameter() *string
	SetNetworkReachableAnalysisId(v string) *GetNetworkReachableAnalysisResponseBody
	GetNetworkReachableAnalysisId() *string
	SetNetworkReachableAnalysisResult(v string) *GetNetworkReachableAnalysisResponseBody
	GetNetworkReachableAnalysisResult() *string
	SetNetworkReachableAnalysisStatus(v string) *GetNetworkReachableAnalysisResponseBody
	GetNetworkReachableAnalysisStatus() *string
	SetReachable(v bool) *GetNetworkReachableAnalysisResponseBody
	GetReachable() *bool
	SetRequestId(v string) *GetNetworkReachableAnalysisResponseBody
	GetRequestId() *string
}

type GetNetworkReachableAnalysisResponseBody struct {
	// The unique ID (UID) of the Alibaba Cloud account.
	//
	// example:
	//
	// 123147627844****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The time when the network path was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-03-16T07:11:27Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The network path ID.
	//
	// example:
	//
	// np-2a1332214fa346b6****
	NetworkPathId *string `json:"NetworkPathId,omitempty" xml:"NetworkPathId,omitempty"`
	// The parameters of the network path.
	//
	// example:
	//
	// {
	//
	//   "sourceId": "i-bp100g5pbp6kj4p9****",
	//
	//   "sourceType": "ecs",
	//
	//   "targetId": "i-t4n4ltwgbbomzb0g****",
	//
	//   "targetType": "ecs"
	//
	// }
	NetworkPathParameter *string `json:"NetworkPathParameter,omitempty" xml:"NetworkPathParameter,omitempty"`
	// The ID of the task for analyzing network reachability.
	//
	// example:
	//
	// nra-8607514e71c1484****
	NetworkReachableAnalysisId *string `json:"NetworkReachableAnalysisId,omitempty" xml:"NetworkReachableAnalysisId,omitempty"`
	// The result of network reachability analysis, which includes the network topology, error codes of network unreachability, and rules of network unreachability.
	//
	// example:
	//
	// {
	//
	//   "errorCode": "",
	//
	//   "networkAclData": {
	//
	//     "networkAclItems": [
	//
	//
	//
	//     ]
	//
	//   },
	//
	//   "nraId": "nra-f2c8701a36424094****",
	//
	//   "requestId": "B931F8A0-620E-5230-B77F-3BD7F612****",
	//
	//   "routeData": {
	//
	//     "routeItems": [
	//
	//
	//
	//     ]
	//
	//   },
	//
	//   "securityGroupData": {
	//
	//     "policy": "accept",
	//
	//     "securityGroupItems": [
	//
	//       {
	//
	//         "description": "default_sg_access_rule",
	//
	//         "matchedRule": {
	//
	//           "bizProtocol": "ALL",
	//
	//           "creatingTime": "2022-11-10T03:24:49Z",
	//
	//           "description": "",
	//
	//           "destinationCidr": "",
	//
	//           "destinationGroupId": "sg-wz980j96p8y99co5****",
	//
	//           "direction": "egress",
	//
	//           "policy": "Accept",
	//
	//           "portRange": "-1/-1",
	//
	//           "priority": "1",
	//
	//           "sourceCidr": "",
	//
	//           "sourceGroupId": ""
	//
	//         },
	//
	//         "policy": "accept",
	//
	//         "resourceId": "eni-wz92ce4saz1jzazg****",
	//
	//         "securityGroupId": "sg-wz980j96p8y99co5****"
	//
	//       },
	//
	//       {
	//
	//         "description": "user_acl_drop_rule",
	//
	//         "matchedRule": {
	//
	//           "bizProtocol": "",
	//
	//           "creatingTime": "",
	//
	//           "description": "",
	//
	//           "destinationCidr": "",
	//
	//           "destinationGroupId": "",
	//
	//           "direction": "",
	//
	//           "policy": "",
	//
	//           "portRange": "",
	//
	//           "priority": "",
	//
	//           "sourceCidr": "",
	//
	//           "sourceGroupId": ""
	//
	//         },
	//
	//         "policy": "",
	//
	//         "resourceId": "eni-wz97vry93t6z4lbd****",
	//
	//         "securityGroupId": "sg-wz980j96p8y99co****"
	//
	//       }
	//
	//     ],
	//
	//     "securityGroupReportId": "sgr-4479d23bb37241aab****"
	//
	//   },
	//
	//   "status": "security_group_checking_target",
	//
	//   "topologyData": {
	//
	//     "positive": {
	//
	//       "linkList": [
	//
	//         {
	//
	//           "id": "i-wz91dk7bor557hp93zyv-->eni-wz92ce4saz1jzazg****",
	//
	//           "source": "i-wz91dk7bor557hp9****",
	//
	//           "target": "eni-wz92ce4saz1jzazg****"
	//
	//         },
	//
	//         {
	//
	//           "id": "eni-wz92ce4saz1jzazgi13d-->vsw-wz9slpwdcppwfrnee****",
	//
	//           "source": "eni-wz92ce4saz1jzazg****",
	//
	//           "target": "vsw-wz9slpwdcppwfrnee****"
	//
	//         },
	//
	//         {
	//
	//           "id": "vsw-wz9slpwdcppwfrneebcrp-->eni-wz97vry93t6z4lbd****",
	//
	//           "source": "vsw-wz9slpwdcppwfrnee****",
	//
	//           "target": "eni-wz97vry93t6z4lbd****"
	//
	//         },
	//
	//         {
	//
	//           "id": "eni-wz97vry93t6z4lbdgmfi-->i-wz91dk7bor557hp9****",
	//
	//           "source": "eni-wz97vry93t6z4lbd****",
	//
	//           "target": "i-wz91dk7bor557hp9****"
	//
	//         }
	//
	//       ],
	//
	//       "nodeList": [
	//
	//         {
	//
	//           "aZone": "cn-shenzhen-d",
	//
	//           "bizInsId": "i-wz91dk7bor557hp9****",
	//
	//           "id": "i-wz91dk7bor557hp9****",
	//
	//           "level": 1,
	//
	//           "matchedRoute": {
	//
	//             "nextHopSet": [
	//
	//
	//
	//             ]
	//
	//           },
	//
	//           "nodeType": "VM",
	//
	//           "regionNo": "cn-shenzhen-st3-a01",
	//
	//           "regionNoAlias": "cn-shenzhen",
	//
	//           "trafficLogs": [
	//
	//
	//
	//           ]
	//
	//         },
	//
	//         {
	//
	//           "aZone": "cn-shenzhen-d",
	//
	//           "bizInsId": "i-wz91dk7bor557hp9****",
	//
	//           "id": "i-wz91dk7bor557hp9****",
	//
	//           "level": 3,
	//
	//           "matchedRoute": {
	//
	//             "nextHopSet": [
	//
	//
	//
	//             ]
	//
	//           },
	//
	//           "nodeType": "VM",
	//
	//           "regionNo": "cn-shenzhen-st3-a01",
	//
	//           "regionNoAlias": "cn-shenzhen",
	//
	//           "trafficLogs": [
	//
	//
	//
	//           ]
	//
	//         },
	//
	//         {
	//
	//           "aZone": "cn-shenzhen-d",
	//
	//           "bizInsId": "vsw-wz9slpwdcppwfrnee****",
	//
	//           "cidr": "192.168.0.0/24",
	//
	//           "id": "vsw-wz9slpwdcppwfrnee****",
	//
	//           "level": 2,
	//
	//           "matchedRoute": {
	//
	//             "nextHopSet": [
	//
	//
	//
	//             ]
	//
	//           },
	//
	//           "nodeType": "VSW",
	//
	//           "regionNo": "cn-shenzhen-st3-a01",
	//
	//           "regionNoAlias": "cn-shenzhen",
	//
	//           "trafficLogs": [
	//
	//
	//
	//           ]
	//
	//         },
	//
	//         {
	//
	//           "bizInsId": "eni-wz92ce4saz1jzazg****",
	//
	//           "id": "eni-wz92ce4saz1jzazg****",
	//
	//           "ip": "192.168.0.33",
	//
	//           "mac": "00:XXXX:3e:16:7c:50",
	//
	//           "matchedRoute": {
	//
	//             "nextHopSet": [
	//
	//
	//
	//             ]
	//
	//           },
	//
	//           "nodeType": "ENI",
	//
	//           "regionNo": "cn-shenzhen-st3-a01",
	//
	//           "regionNoAlias": "cn-shenzhen",
	//
	//           "status": "InUse",
	//
	//           "trafficLogs": [
	//
	//
	//
	//           ]
	//
	//         },
	//
	//         {
	//
	//           "bizInsId": "eni-wz97vry93t6z4lbd****",
	//
	//           "id": "eni-wz97vry93t6z4lbd****",
	//
	//           "ip": "192.168.0.34",
	//
	//           "mac": "00:XXXX:3e:14:70:c2",
	//
	//           "matchedRoute": {
	//
	//             "nextHopSet": [
	//
	//
	//
	//             ]
	//
	//           },
	//
	//           "nodeType": "ENI",
	//
	//           "regionNo": "cn-shenzhen-st3-a01",
	//
	//           "regionNoAlias": "cn-shenzhen",
	//
	//           "status": "InUse",
	//
	//           "trafficLogs": [
	//
	//
	//
	//           ]
	//
	//         }
	//
	//       ]
	//
	//     },
	//
	//     "reverse": {
	//
	//       "revLinkList": [
	//
	//         {
	//
	//           "id": "i-wz91dk7bor557hp93zys-->eni-wz97vry93t6z4lbd****",
	//
	//           "source": "i-wz91dk7bor557hp9****",
	//
	//           "target": "eni-wz97vry93t6z4lbd****"
	//
	//         },
	//
	//         {
	//
	//           "id": "eni-wz97vry93t6z4lbdgmfi-->vsw-wz9slpwdcppwfrnee****",
	//
	//           "source": "eni-wz97vry93t6z4lbd****",
	//
	//           "target": "vsw-wz9slpwdcppwfrnee****"
	//
	//         },
	//
	//         {
	//
	//           "id": "vsw-wz9slpwdcppwfrneebcrp-->eni-wz92ce4saz1jzazg****",
	//
	//           "source": "vsw-wz9slpwdcppwfrnee****",
	//
	//           "target": "eni-wz92ce4saz1jzazg****"
	//
	//         },
	//
	//         {
	//
	//           "id": "eni-wz92ce4saz1jzazgi13d-->i-wz91dk7bor557hp9****",
	//
	//           "source": "eni-wz92ce4saz1jzazg****",
	//
	//           "target": "i-wz91dk7bor557hp9****"
	//
	//         }
	//
	//       ],
	//
	//       "revNodeList": [
	//
	//         {
	//
	//           "aZone": "cn-shenzhen-d",
	//
	//           "bizInsId": "i-wz91dk7bor557hp9****",
	//
	//           "id": "i-wz91dk7bor557hp9****",
	//
	//           "level": 1,
	//
	//           "nodeType": "VM",
	//
	//           "regionNo": "cn-shenzhen-st3-a01",
	//
	//           "regionNoAlias": "cn-shenzhen",
	//
	//           "revMatchedRoute": {
	//
	//             "revNextHopSet": [
	//
	//
	//
	//             ]
	//
	//           }
	//
	//         },
	//
	//         {
	//
	//           "aZone": "cn-shenzhen-d",
	//
	//           "bizInsId": "i-wz91dk7bor557hp9****",
	//
	//           "id": "i-wz91dk7bor557hp9****",
	//
	//           "level": 3,
	//
	//           "nodeType": "VM",
	//
	//           "regionNo": "cn-shenzhen-st3-a01",
	//
	//           "regionNoAlias": "cn-shenzhen",
	//
	//           "revMatchedRoute": {
	//
	//             "revNextHopSet": [
	//
	//
	//
	//             ]
	//
	//           }
	//
	//         },
	//
	//         {
	//
	//           "aZone": "cn-shenzhen-d",
	//
	//           "bizInsId": "vsw-wz9slpwdcppwfrnee****",
	//
	//           "cidr": "192.168.0.0/24",
	//
	//           "id": "vsw-wz9slpwdcppwfrnee****",
	//
	//           "level": 2,
	//
	//           "nodeType": "VSW",
	//
	//           "regionNo": "cn-shenzhen-st3-a01",
	//
	//           "regionNoAlias": "cn-shenzhen",
	//
	//           "revMatchedRoute": {
	//
	//             "revNextHopSet": [
	//
	//
	//
	//             ]
	//
	//           }
	//
	//         },
	//
	//         {
	//
	//           "bizInsId": "eni-wz97vry93t6z4lbd****",
	//
	//           "id": "eni-wz97vry93t6z4lbd****",
	//
	//           "ip": "192.168.0.34",
	//
	//           "mac": "00:XXXX:3e:14:70:c2",
	//
	//           "nodeType": "ENI",
	//
	//           "regionNo": "cn-shenzhen-st3-a01",
	//
	//           "regionNoAlias": "cn-shenzhen",
	//
	//           "revMatchedRoute": {
	//
	//             "revNextHopSet": [
	//
	//
	//
	//             ]
	//
	//           },
	//
	//           "status": "InUse"
	//
	//         },
	//
	//         {
	//
	//           "bizInsId": "eni-wz92ce4saz1jzazg****",
	//
	//           "id": "eni-wz92ce4saz1jzazg****",
	//
	//           "ip": "192.168.0.33",
	//
	//           "mac": "00:XXXX:3e:16:7c:50",
	//
	//           "nodeType": "ENI",
	//
	//           "regionNo": "cn-shenzhen-st3-a01",
	//
	//           "regionNoAlias": "cn-shenzhen",
	//
	//           "revMatchedRoute": {
	//
	//             "revNextHopSet": [
	//
	//
	//
	//             ]
	//
	//           },
	//
	//           "status": "InUse"
	//
	//         }
	//
	//       ]
	//
	//     },
	//
	//     "topologyReportId": "tpr-21cf60002715491b8****"
	//
	//   }
	//
	// }
	NetworkReachableAnalysisResult *string `json:"NetworkReachableAnalysisResult,omitempty" xml:"NetworkReachableAnalysisResult,omitempty"`
	// The state of the task for analyzing network reachability. Valid values:
	//
	// 	- **init**: The task is in progress.
	//
	// 	- **finish**: The task is complete.
	//
	// 	- **error**: An analysis error occurred.
	//
	// 	- **timeout**: The task timed out.
	//
	// example:
	//
	// finish
	NetworkReachableAnalysisStatus *string `json:"NetworkReachableAnalysisStatus,omitempty" xml:"NetworkReachableAnalysisStatus,omitempty"`
	// Indicates whether the network path is reachable. Valid values:
	//
	// 	- **true**: The network path is reachable.
	//
	// 	- **false**: The network path is unreachable.
	//
	// example:
	//
	// true
	Reachable *bool `json:"Reachable,omitempty" xml:"Reachable,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DEE0FEAF-59AE-5CDD-AA07-626BC365D571
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNetworkReachableAnalysisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkReachableAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *GetNetworkReachableAnalysisResponseBody) GetAliUid() *int64 {
	return s.AliUid
}

func (s *GetNetworkReachableAnalysisResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetNetworkReachableAnalysisResponseBody) GetNetworkPathId() *string {
	return s.NetworkPathId
}

func (s *GetNetworkReachableAnalysisResponseBody) GetNetworkPathParameter() *string {
	return s.NetworkPathParameter
}

func (s *GetNetworkReachableAnalysisResponseBody) GetNetworkReachableAnalysisId() *string {
	return s.NetworkReachableAnalysisId
}

func (s *GetNetworkReachableAnalysisResponseBody) GetNetworkReachableAnalysisResult() *string {
	return s.NetworkReachableAnalysisResult
}

func (s *GetNetworkReachableAnalysisResponseBody) GetNetworkReachableAnalysisStatus() *string {
	return s.NetworkReachableAnalysisStatus
}

func (s *GetNetworkReachableAnalysisResponseBody) GetReachable() *bool {
	return s.Reachable
}

func (s *GetNetworkReachableAnalysisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNetworkReachableAnalysisResponseBody) SetAliUid(v int64) *GetNetworkReachableAnalysisResponseBody {
	s.AliUid = &v
	return s
}

func (s *GetNetworkReachableAnalysisResponseBody) SetCreateTime(v string) *GetNetworkReachableAnalysisResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetNetworkReachableAnalysisResponseBody) SetNetworkPathId(v string) *GetNetworkReachableAnalysisResponseBody {
	s.NetworkPathId = &v
	return s
}

func (s *GetNetworkReachableAnalysisResponseBody) SetNetworkPathParameter(v string) *GetNetworkReachableAnalysisResponseBody {
	s.NetworkPathParameter = &v
	return s
}

func (s *GetNetworkReachableAnalysisResponseBody) SetNetworkReachableAnalysisId(v string) *GetNetworkReachableAnalysisResponseBody {
	s.NetworkReachableAnalysisId = &v
	return s
}

func (s *GetNetworkReachableAnalysisResponseBody) SetNetworkReachableAnalysisResult(v string) *GetNetworkReachableAnalysisResponseBody {
	s.NetworkReachableAnalysisResult = &v
	return s
}

func (s *GetNetworkReachableAnalysisResponseBody) SetNetworkReachableAnalysisStatus(v string) *GetNetworkReachableAnalysisResponseBody {
	s.NetworkReachableAnalysisStatus = &v
	return s
}

func (s *GetNetworkReachableAnalysisResponseBody) SetReachable(v bool) *GetNetworkReachableAnalysisResponseBody {
	s.Reachable = &v
	return s
}

func (s *GetNetworkReachableAnalysisResponseBody) SetRequestId(v string) *GetNetworkReachableAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNetworkReachableAnalysisResponseBody) Validate() error {
	return dara.Validate(s)
}
