// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCampaignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCampaignId(v string) *CreateCampaignResponseBody
	GetCampaignId() *string
	SetCode(v string) *CreateCampaignResponseBody
	GetCode() *string
	SetHttpStatusCode(v int64) *CreateCampaignResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *CreateCampaignResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateCampaignResponseBody
	GetRequestId() *string
}

type CreateCampaignResponseBody struct {
	CampaignId     *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCampaignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCampaignResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCampaignResponseBody) GetCampaignId() *string {
	return s.CampaignId
}

func (s *CreateCampaignResponseBody) GetCode() *string {
	return s.Code
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

func (s *CreateCampaignResponseBody) SetCampaignId(v string) *CreateCampaignResponseBody {
	s.CampaignId = &v
	return s
}

func (s *CreateCampaignResponseBody) SetCode(v string) *CreateCampaignResponseBody {
	s.Code = &v
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
