// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCampaignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateCampaignResponseBody
	GetCode() *string
	SetData(v string) *CreateCampaignResponseBody
	GetData() *string
	SetHttpStatusCode(v int64) *CreateCampaignResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *CreateCampaignResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateCampaignResponseBody
	GetRequestId() *string
}

type CreateCampaignResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// c58b9719-3bc3-441d-a4d3-fc0309ef7066
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7CC6523B-0E51-1B62-8DA5-6A9831CAE315
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCampaignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCampaignResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCampaignResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCampaignResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateCampaignResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *CreateCampaignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCampaignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCampaignResponseBody) SetCode(v string) *CreateCampaignResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCampaignResponseBody) SetData(v string) *CreateCampaignResponseBody {
	s.Data = &v
	return s
}

func (s *CreateCampaignResponseBody) SetHttpStatusCode(v int64) *CreateCampaignResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateCampaignResponseBody) SetMessage(v string) *CreateCampaignResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCampaignResponseBody) SetRequestId(v string) *CreateCampaignResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCampaignResponseBody) Validate() error {
	return dara.Validate(s)
}
