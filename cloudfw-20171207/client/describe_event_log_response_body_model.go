// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeEventLogResponseBodyDataList) *DescribeEventLogResponseBody
	GetDataList() []*DescribeEventLogResponseBodyDataList
	SetRequestId(v string) *DescribeEventLogResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeEventLogResponseBody
	GetTotalCount() *int64
}

type DescribeEventLogResponseBody struct {
	// The list of ISPs.
	DataList []*DescribeEventLogResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D19D8F70-D64B-5A95-905A-6073BF4A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 12
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEventLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEventLogResponseBody) GetDataList() []*DescribeEventLogResponseBodyDataList {
	return s.DataList
}

func (s *DescribeEventLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEventLogResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeEventLogResponseBody) SetDataList(v []*DescribeEventLogResponseBodyDataList) *DescribeEventLogResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeEventLogResponseBody) SetRequestId(v string) *DescribeEventLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEventLogResponseBody) SetTotalCount(v int64) *DescribeEventLogResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeEventLogResponseBody) Validate() error {
	if s.DataList != nil {
		for _, item := range s.DataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEventLogResponseBodyDataList struct {
	// The application type supported by the access control policy. We recommend that you use ApplicationNameList. Valid values:
	//
	// - **FTP**
	//
	// - **HTTP**
	//
	// - **HTTPS**
	//
	// - **Memcache**
	//
	// - **MongoDB**
	//
	// - **MQTT**
	//
	// - **MySQL**
	//
	// - **RDP**
	//
	// - **Redis**
	//
	// - **SMTP**
	//
	// - **SMTPS**
	//
	// - **SSH**
	//
	// - **SSL**
	//
	// - **VNC**
	//
	// - **ANY*	- (all application types)
	//
	// example:
	//
	// ANY
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The name of the attacked application.
	//
	// example:
	//
	// MySql
	AttackApp *string `json:"AttackApp,omitempty" xml:"AttackApp,omitempty"`
	// The attack type of the intrusion prevention event. Valid values:
	//
	// - **1**: Abnormal connection.
	//
	// - **2**: Command execution.
	//
	// - **3**: Brute-force attack.
	//
	// - **4**: Scanning.
	//
	// - **5**: Other.
	//
	// - **6**: Information leakage.
	//
	// - **7**: DoS attack.
	//
	// - **8**: Overflow attack.
	//
	// - **9**: Web attack.
	//
	// - **10**: Trojan backdoor.
	//
	// - **11**: Viruses and worms.
	//
	// - **12**: Mining behavior.
	//
	// - **13**: Reverse shell.
	//
	// example:
	//
	// 0
	AttackType *int64 `json:"AttackType,omitempty" xml:"AttackType,omitempty"`
	// The traffic direction of the intrusion prevention event. Valid values:
	//
	// - **in**: Inbound.
	//
	// - **out**: Outbound.
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The destination IP address that is queried. This indicates that the intrusion prevention event involves this destination IP address.
	//
	// example:
	//
	// 2.2.2.2
	DstIP *string `json:"DstIP,omitempty" xml:"DstIP,omitempty"`
	// The destination port.
	//
	// example:
	//
	// 80
	DstPort *int32 `json:"DstPort,omitempty" xml:"DstPort,omitempty"`
	// The hit content.
	//
	// example:
	//
	// www.example.com
	HitContent *string `json:"HitContent,omitempty" xml:"HitContent,omitempty"`
	// The hit type.
	//
	// example:
	//
	// domain
	HitType *string `json:"HitType,omitempty" xml:"HitType,omitempty"`
	// The protocol type.
	//
	// example:
	//
	// tcp
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The UID of the Cloud Firewall member account.
	//
	// example:
	//
	// 14151892****7022
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The ID of the protection rule for the intrusion prevention event.
	//
	// example:
	//
	// 00000000-0000-0000-0000-000000000000
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The rule name.
	//
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The defense status. Valid values:
	//
	// - **1**: Alert.
	//
	// - **2**: Block.
	//
	// example:
	//
	// 0
	RuleResult *int32 `json:"RuleResult,omitempty" xml:"RuleResult,omitempty"`
	// The source of the detection rule for the intrusion prevention event. Valid values:
	//
	// - **1**: Basic protection.
	//
	// - **2**: Virtual patches.
	//
	// - **4**: Threat intelligence.
	//
	// example:
	//
	// 0
	RuleSource *int64 `json:"RuleSource,omitempty" xml:"RuleSource,omitempty"`
	// The SNI used when the TLS connection is established.
	//
	// example:
	//
	// www.example.com
	Sni *string `json:"Sni,omitempty" xml:"Sni,omitempty"`
	// The source IP address that is queried. This indicates that the intrusion prevention event involves this source IP address.
	//
	// example:
	//
	// 192.0.XX.XX
	SrcIP *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	// The source port.
	//
	// example:
	//
	// 20206
	SrcPort *int32 `json:"SrcPort,omitempty" xml:"SrcPort,omitempty"`
	// The cipher suite used for TLS handshake negotiation, such as TLS_AES_128_GCM_SHA256.
	//
	// example:
	//
	// TLS_AES_128_GCM_SHA256
	SslCipher *string `json:"SslCipher,omitempty" xml:"SslCipher,omitempty"`
	// The phase in which the TLS handshake fails, such as Certificate.
	//
	// example:
	//
	// Certificate
	SslErrorIndex *string `json:"SslErrorIndex,omitempty" xml:"SslErrorIndex,omitempty"`
	// The reason why the TLS handshake fails, such as BAD_PACKET_LENGTH.
	//
	// example:
	//
	// BAD_PACKET_LENGTH
	SslErrorMsg *string `json:"SslErrorMsg,omitempty" xml:"SslErrorMsg,omitempty"`
	// The TLS handshake result. For example, success indicates that the handshake is successful.
	//
	// example:
	//
	// success
	SslHandshakeRes *string `json:"SslHandshakeRes,omitempty" xml:"SslHandshakeRes,omitempty"`
	// The TLS version, such as TLSv1.3.
	//
	// example:
	//
	// TLSv1.3
	SslProtocol *string `json:"SslProtocol,omitempty" xml:"SslProtocol,omitempty"`
	// The start time of the data. The value is a timestamp in seconds.
	//
	// example:
	//
	// 1751423362
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The document path.
	//
	// example:
	//
	// http://example.com/task3%DF\\"
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The severity level of the intrusion prevention event. Valid values:
	//
	// - **1**: Low.
	//
	// - **2**: Medium.
	//
	// - **3**: High.
	//
	// example:
	//
	// 0
	VulLevel *int64 `json:"VulLevel,omitempty" xml:"VulLevel,omitempty"`
}

func (s DescribeEventLogResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventLogResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeEventLogResponseBodyDataList) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *DescribeEventLogResponseBodyDataList) GetAttackApp() *string {
	return s.AttackApp
}

func (s *DescribeEventLogResponseBodyDataList) GetAttackType() *int64 {
	return s.AttackType
}

func (s *DescribeEventLogResponseBodyDataList) GetDirection() *string {
	return s.Direction
}

func (s *DescribeEventLogResponseBodyDataList) GetDstIP() *string {
	return s.DstIP
}

func (s *DescribeEventLogResponseBodyDataList) GetDstPort() *int32 {
	return s.DstPort
}

func (s *DescribeEventLogResponseBodyDataList) GetHitContent() *string {
	return s.HitContent
}

func (s *DescribeEventLogResponseBodyDataList) GetHitType() *string {
	return s.HitType
}

func (s *DescribeEventLogResponseBodyDataList) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *DescribeEventLogResponseBodyDataList) GetMemberUid() *int64 {
	return s.MemberUid
}

func (s *DescribeEventLogResponseBodyDataList) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeEventLogResponseBodyDataList) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeEventLogResponseBodyDataList) GetRuleResult() *int32 {
	return s.RuleResult
}

func (s *DescribeEventLogResponseBodyDataList) GetRuleSource() *int64 {
	return s.RuleSource
}

func (s *DescribeEventLogResponseBodyDataList) GetSni() *string {
	return s.Sni
}

func (s *DescribeEventLogResponseBodyDataList) GetSrcIP() *string {
	return s.SrcIP
}

func (s *DescribeEventLogResponseBodyDataList) GetSrcPort() *int32 {
	return s.SrcPort
}

func (s *DescribeEventLogResponseBodyDataList) GetSslCipher() *string {
	return s.SslCipher
}

func (s *DescribeEventLogResponseBodyDataList) GetSslErrorIndex() *string {
	return s.SslErrorIndex
}

func (s *DescribeEventLogResponseBodyDataList) GetSslErrorMsg() *string {
	return s.SslErrorMsg
}

func (s *DescribeEventLogResponseBodyDataList) GetSslHandshakeRes() *string {
	return s.SslHandshakeRes
}

func (s *DescribeEventLogResponseBodyDataList) GetSslProtocol() *string {
	return s.SslProtocol
}

func (s *DescribeEventLogResponseBodyDataList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeEventLogResponseBodyDataList) GetUrl() *string {
	return s.Url
}

func (s *DescribeEventLogResponseBodyDataList) GetVulLevel() *int64 {
	return s.VulLevel
}

func (s *DescribeEventLogResponseBodyDataList) SetApplicationName(v string) *DescribeEventLogResponseBodyDataList {
	s.ApplicationName = &v
	return s
}

func (s *DescribeEventLogResponseBodyDataList) SetAttackApp(v string) *DescribeEventLogResponseBodyDataList {
	s.AttackApp = &v
	return s
}

func (s *DescribeEventLogResponseBodyDataList) SetAttackType(v int64) *DescribeEventLogResponseBodyDataList {
	s.AttackType = &v
	return s
}

func (s *DescribeEventLogResponseBodyDataList) SetDirection(v string) *DescribeEventLogResponseBodyDataList {
	s.Direction = &v
	return s
}

func (s *DescribeEventLogResponseBodyDataList) SetDstIP(v string) *DescribeEventLogResponseBodyDataList {
	s.DstIP = &v
	return s
}

func (s *DescribeEventLogResponseBodyDataList) SetDstPort(v int32) *DescribeEventLogResponseBodyDataList {
	s.DstPort = &v
	return s
}

func (s *DescribeEventLogResponseBodyDataList) SetHitContent(v string) *DescribeEventLogResponseBodyDataList {
	s.HitContent = &v
	return s
}

func (s *DescribeEventLogResponseBodyDataList) SetHitType(v string) *DescribeEventLogResponseBodyDataList {
	s.HitType = &v
	return s
}

func (s *DescribeEventLogResponseBodyDataList) SetIpProtocol(v string) *DescribeEventLogResponseBodyDataList {
	s.IpProtocol = &v
	return s
}

func (s *DescribeEventLogResponseBodyDataList) SetMemberUid(v int64) *DescribeEventLogResponseBodyDataList {
	s.MemberUid = &v
	return s
}

func (s *DescribeEventLogResponseBodyDataList) SetRuleId(v string) *DescribeEventLogResponseBodyDataList {
	s.RuleId = &v
	return s
}

func (s *DescribeEventLogResponseBodyDataList) SetRuleName(v string) *DescribeEventLogResponseBodyDataList {
	s.RuleName = &v
	return s
}

func (s *DescribeEventLogResponseBodyDataList) SetRuleResult(v int32) *DescribeEventLogResponseBodyDataList {
	s.RuleResult = &v
	return s
}

func (s *DescribeEventLogResponseBodyDataList) SetRuleSource(v int64) *DescribeEventLogResponseBodyDataList {
	s.RuleSource = &v
	return s
}

func (s *DescribeEventLogResponseBodyDataList) SetSni(v string) *DescribeEventLogResponseBodyDataList {
	s.Sni = &v
	return s
}

func (s *DescribeEventLogResponseBodyDataList) SetSrcIP(v string) *DescribeEventLogResponseBodyDataList {
	s.SrcIP = &v
	return s
}

func (s *DescribeEventLogResponseBodyDataList) SetSrcPort(v int32) *DescribeEventLogResponseBodyDataList {
	s.SrcPort = &v
	return s
}

func (s *DescribeEventLogResponseBodyDataList) SetSslCipher(v string) *DescribeEventLogResponseBodyDataList {
	s.SslCipher = &v
	return s
}

func (s *DescribeEventLogResponseBodyDataList) SetSslErrorIndex(v string) *DescribeEventLogResponseBodyDataList {
	s.SslErrorIndex = &v
	return s
}

func (s *DescribeEventLogResponseBodyDataList) SetSslErrorMsg(v string) *DescribeEventLogResponseBodyDataList {
	s.SslErrorMsg = &v
	return s
}

func (s *DescribeEventLogResponseBodyDataList) SetSslHandshakeRes(v string) *DescribeEventLogResponseBodyDataList {
	s.SslHandshakeRes = &v
	return s
}

func (s *DescribeEventLogResponseBodyDataList) SetSslProtocol(v string) *DescribeEventLogResponseBodyDataList {
	s.SslProtocol = &v
	return s
}

func (s *DescribeEventLogResponseBodyDataList) SetStartTime(v int64) *DescribeEventLogResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *DescribeEventLogResponseBodyDataList) SetUrl(v string) *DescribeEventLogResponseBodyDataList {
	s.Url = &v
	return s
}

func (s *DescribeEventLogResponseBodyDataList) SetVulLevel(v int64) *DescribeEventLogResponseBodyDataList {
	s.VulLevel = &v
	return s
}

func (s *DescribeEventLogResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
