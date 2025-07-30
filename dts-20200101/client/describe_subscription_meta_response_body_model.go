// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubscriptionMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DescribeSubscriptionMetaResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeSubscriptionMetaResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v string) *DescribeSubscriptionMetaResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *DescribeSubscriptionMetaResponseBody
	GetRequestId() *string
	SetSubscriptionMetaList(v []*DescribeSubscriptionMetaResponseBodySubscriptionMetaList) *DescribeSubscriptionMetaResponseBody
	GetSubscriptionMetaList() []*DescribeSubscriptionMetaResponseBodySubscriptionMetaList
	SetSuccess(v string) *DescribeSubscriptionMetaResponseBody
	GetSuccess() *string
}

type DescribeSubscriptionMetaResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// The Value of Input Parameter %s is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C12E7A51-09A4-5796-94BE-08B6DA******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the subtasks.
	SubscriptionMetaList []*DescribeSubscriptionMetaResponseBodySubscriptionMetaList `json:"SubscriptionMetaList,omitempty" xml:"SubscriptionMetaList,omitempty" type:"Repeated"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSubscriptionMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionMetaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionMetaResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeSubscriptionMetaResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeSubscriptionMetaResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeSubscriptionMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSubscriptionMetaResponseBody) GetSubscriptionMetaList() []*DescribeSubscriptionMetaResponseBodySubscriptionMetaList {
	return s.SubscriptionMetaList
}

func (s *DescribeSubscriptionMetaResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSubscriptionMetaResponseBody) SetErrCode(v string) *DescribeSubscriptionMetaResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSubscriptionMetaResponseBody) SetErrMessage(v string) *DescribeSubscriptionMetaResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSubscriptionMetaResponseBody) SetHttpStatusCode(v string) *DescribeSubscriptionMetaResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeSubscriptionMetaResponseBody) SetRequestId(v string) *DescribeSubscriptionMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSubscriptionMetaResponseBody) SetSubscriptionMetaList(v []*DescribeSubscriptionMetaResponseBodySubscriptionMetaList) *DescribeSubscriptionMetaResponseBody {
	s.SubscriptionMetaList = v
	return s
}

func (s *DescribeSubscriptionMetaResponseBody) SetSuccess(v string) *DescribeSubscriptionMetaResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSubscriptionMetaResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSubscriptionMetaResponseBodySubscriptionMetaList struct {
	// The consumer offset of the subtask. It is a UNIX timestamp that is generated when the client consumes the first data record. Unit: seconds.
	//
	// >  You can use a search engine to obtain a UNIX timestamp converter.
	//
	// example:
	//
	// 1610524452
	Checkpoint *int64 `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	// The objects of the subtask. For more information, see [Objects of DTS tasks](https://help.aliyun.com/document_detail/209545.html).
	//
	// example:
	//
	// {\\"dtstestdata\\":{\\"name\\":\\"dtstestdata\\",\\"all\\":false,\\"Table\\":{\\"order\\":{\\"name\\":\\"order\\",\\"all\\":true}}}}
	DBList *string `json:"DBList,omitempty" xml:"DBList,omitempty"`
	// The endpoint and port number of the change tracking instance.
	//
	// example:
	//
	// dts-cn-hangzhou.aliyuncs.com:18001
	DProxyUrl *string `json:"DProxyUrl,omitempty" xml:"DProxyUrl,omitempty"`
	// The consumer group ID of the subtask.
	//
	// example:
	//
	// z38m91gg2******
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	// The topic of the subtask.
	//
	// example:
	//
	// cn_hangzhou_rm_bp1n0x0x5tz******_dtstestdata_version2
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s DescribeSubscriptionMetaResponseBodySubscriptionMetaList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionMetaResponseBodySubscriptionMetaList) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionMetaResponseBodySubscriptionMetaList) GetCheckpoint() *int64 {
	return s.Checkpoint
}

func (s *DescribeSubscriptionMetaResponseBodySubscriptionMetaList) GetDBList() *string {
	return s.DBList
}

func (s *DescribeSubscriptionMetaResponseBodySubscriptionMetaList) GetDProxyUrl() *string {
	return s.DProxyUrl
}

func (s *DescribeSubscriptionMetaResponseBodySubscriptionMetaList) GetSid() *string {
	return s.Sid
}

func (s *DescribeSubscriptionMetaResponseBodySubscriptionMetaList) GetTopic() *string {
	return s.Topic
}

func (s *DescribeSubscriptionMetaResponseBodySubscriptionMetaList) SetCheckpoint(v int64) *DescribeSubscriptionMetaResponseBodySubscriptionMetaList {
	s.Checkpoint = &v
	return s
}

func (s *DescribeSubscriptionMetaResponseBodySubscriptionMetaList) SetDBList(v string) *DescribeSubscriptionMetaResponseBodySubscriptionMetaList {
	s.DBList = &v
	return s
}

func (s *DescribeSubscriptionMetaResponseBodySubscriptionMetaList) SetDProxyUrl(v string) *DescribeSubscriptionMetaResponseBodySubscriptionMetaList {
	s.DProxyUrl = &v
	return s
}

func (s *DescribeSubscriptionMetaResponseBodySubscriptionMetaList) SetSid(v string) *DescribeSubscriptionMetaResponseBodySubscriptionMetaList {
	s.Sid = &v
	return s
}

func (s *DescribeSubscriptionMetaResponseBodySubscriptionMetaList) SetTopic(v string) *DescribeSubscriptionMetaResponseBodySubscriptionMetaList {
	s.Topic = &v
	return s
}

func (s *DescribeSubscriptionMetaResponseBodySubscriptionMetaList) Validate() error {
	return dara.Validate(s)
}
