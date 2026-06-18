// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineRuntimeInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHotlineRuntimeInfoResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *GetHotlineRuntimeInfoResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *GetHotlineRuntimeInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotlineRuntimeInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetHotlineRuntimeInfoResponseBody
	GetSuccess() *bool
}

type GetHotlineRuntimeInfoResponseBody struct {
	// The status code. A return value of "Success" indicates that the request succeeded.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Hotline runtime information.
	//
	// example:
	//
	// {"buId":905,"phoneToken":"roESVzzgD1ytmG0W6nMkWinI6fbpKovm14fBhA0NsYPyi/srX/G0SvNB2z96AYiqy1vxQHiaFOdZSxtsAubdgr2kjs2yas7COO5dukQpsOLq9iNI4U6sKlcvaBi8xsyUr/hyqCdTVZDcYCOq0lH6eeNIQK/f7/gWqIw****=","isNeedWorkShift":false,"servicerId":741018,"accConfigUrl":"pre-acc-cs-public.alibaba-inc.com","agentBasicCode":"AgentCheckout","startWorkToReady":true,"servicerDn":"2000000001904123","isMaster":"1","tenantId":2001,"depId":1139139,"accTenantId":2001,"phoneData":"{\\"timeStamp\\":164240****,\\"expireTime\\":164257****,\\"dn\\":\\"200000000190****\\",\\"source\\":\\"xspace\\",\\"serviceId\\":741018}","agentBasicDesc":"坐席签出状态"}
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHotlineRuntimeInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineRuntimeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineRuntimeInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHotlineRuntimeInfoResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetHotlineRuntimeInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotlineRuntimeInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotlineRuntimeInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetHotlineRuntimeInfoResponseBody) SetCode(v string) *GetHotlineRuntimeInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineRuntimeInfoResponseBody) SetData(v map[string]interface{}) *GetHotlineRuntimeInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetHotlineRuntimeInfoResponseBody) SetMessage(v string) *GetHotlineRuntimeInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineRuntimeInfoResponseBody) SetRequestId(v string) *GetHotlineRuntimeInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineRuntimeInfoResponseBody) SetSuccess(v bool) *GetHotlineRuntimeInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotlineRuntimeInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
