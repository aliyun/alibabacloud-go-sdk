// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCampaignNumbersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyCampaignNumbersResponseBody
	GetCode() *string
	SetData(v interface{}) *ModifyCampaignNumbersResponseBody
	GetData() interface{}
	SetHttpStatusCode(v int32) *ModifyCampaignNumbersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyCampaignNumbersResponseBody
	GetMessage() *string
	SetParams(v []*string) *ModifyCampaignNumbersResponseBody
	GetParams() []*string
	SetRequestId(v string) *ModifyCampaignNumbersResponseBody
	GetRequestId() *string
}

type ModifyCampaignNumbersResponseBody struct {
	// example:
	//
	// OK
	Code *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// DE803553-8AA9-4B9D-9E4E-A82BC69EDCEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCampaignNumbersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCampaignNumbersResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCampaignNumbersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyCampaignNumbersResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ModifyCampaignNumbersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyCampaignNumbersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyCampaignNumbersResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ModifyCampaignNumbersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCampaignNumbersResponseBody) SetCode(v string) *ModifyCampaignNumbersResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyCampaignNumbersResponseBody) SetData(v interface{}) *ModifyCampaignNumbersResponseBody {
	s.Data = v
	return s
}

func (s *ModifyCampaignNumbersResponseBody) SetHttpStatusCode(v int32) *ModifyCampaignNumbersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyCampaignNumbersResponseBody) SetMessage(v string) *ModifyCampaignNumbersResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyCampaignNumbersResponseBody) SetParams(v []*string) *ModifyCampaignNumbersResponseBody {
	s.Params = v
	return s
}

func (s *ModifyCampaignNumbersResponseBody) SetRequestId(v string) *ModifyCampaignNumbersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCampaignNumbersResponseBody) Validate() error {
	return dara.Validate(s)
}
