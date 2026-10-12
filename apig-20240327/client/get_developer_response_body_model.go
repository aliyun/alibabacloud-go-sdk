// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeveloperResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDeveloperResponseBody
	GetCode() *string
	SetData(v *GetDeveloperResponseBodyData) *GetDeveloperResponseBody
	GetData() *GetDeveloperResponseBodyData
	SetMessage(v string) *GetDeveloperResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDeveloperResponseBody
	GetRequestId() *string
}

type GetDeveloperResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *GetDeveloperResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetDeveloperResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeveloperResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeveloperResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDeveloperResponseBody) GetData() *GetDeveloperResponseBodyData {
	return s.Data
}

func (s *GetDeveloperResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDeveloperResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeveloperResponseBody) SetCode(v string) *GetDeveloperResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeveloperResponseBody) SetData(v *GetDeveloperResponseBodyData) *GetDeveloperResponseBody {
	s.Data = v
	return s
}

func (s *GetDeveloperResponseBody) SetMessage(v string) *GetDeveloperResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeveloperResponseBody) SetRequestId(v string) *GetDeveloperResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeveloperResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDeveloperResponseBodyData struct {
	// The developer ID. This ID can be used to identify the developer in API operations such as SetHiMarketDeveloperStatus and DeleteHiMarketDeveloper.
	//
	// example:
	//
	// dev-xxxxxxxx
	DeveloperId *string `json:"developerId,omitempty" xml:"developerId,omitempty"`
	// The ID of the portal to which the developer belongs. This ID can be used to identify the portal in API operations such as GetPortal, UpdatePortal, and DeletePortal.
	//
	// example:
	//
	// portal-xxxxxxxx
	PortalId *string `json:"portalId,omitempty" xml:"portalId,omitempty"`
	// The developer status. Valid values: PENDING, APPROVED, and REJECTED.
	//
	// example:
	//
	// APPROVED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The logon username of the developer.
	//
	// example:
	//
	// developer-demo
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s GetDeveloperResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDeveloperResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeveloperResponseBodyData) GetDeveloperId() *string {
	return s.DeveloperId
}

func (s *GetDeveloperResponseBodyData) GetPortalId() *string {
	return s.PortalId
}

func (s *GetDeveloperResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetDeveloperResponseBodyData) GetUsername() *string {
	return s.Username
}

func (s *GetDeveloperResponseBodyData) SetDeveloperId(v string) *GetDeveloperResponseBodyData {
	s.DeveloperId = &v
	return s
}

func (s *GetDeveloperResponseBodyData) SetPortalId(v string) *GetDeveloperResponseBodyData {
	s.PortalId = &v
	return s
}

func (s *GetDeveloperResponseBodyData) SetStatus(v string) *GetDeveloperResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetDeveloperResponseBodyData) SetUsername(v string) *GetDeveloperResponseBodyData {
	s.Username = &v
	return s
}

func (s *GetDeveloperResponseBodyData) Validate() error {
	return dara.Validate(s)
}
