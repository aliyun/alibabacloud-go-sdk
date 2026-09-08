// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseCampaignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PauseCampaignResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *PauseCampaignResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *PauseCampaignResponseBody
	GetMessage() *string
	SetRequestId(v string) *PauseCampaignResponseBody
	GetRequestId() *string
}

type PauseCampaignResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// 无
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4A68E287-6888-5ADB-8048-DB488B4DEF35
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PauseCampaignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PauseCampaignResponseBody) GoString() string {
	return s.String()
}

func (s *PauseCampaignResponseBody) GetCode() *string {
	return s.Code
}

func (s *PauseCampaignResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *PauseCampaignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PauseCampaignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PauseCampaignResponseBody) SetCode(v string) *PauseCampaignResponseBody {
	s.Code = &v
	return s
}

func (s *PauseCampaignResponseBody) SetHttpStatusCode(v string) *PauseCampaignResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PauseCampaignResponseBody) SetMessage(v string) *PauseCampaignResponseBody {
	s.Message = &v
	return s
}

func (s *PauseCampaignResponseBody) SetRequestId(v string) *PauseCampaignResponseBody {
	s.RequestId = &v
	return s
}

func (s *PauseCampaignResponseBody) Validate() error {
	return dara.Validate(s)
}
