// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePortalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePortalResponseBody
	GetCode() *string
	SetData(v *CreatePortalResponseBodyData) *CreatePortalResponseBody
	GetData() *CreatePortalResponseBodyData
	SetMessage(v string) *CreatePortalResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreatePortalResponseBody
	GetRequestId() *string
}

type CreatePortalResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *CreatePortalResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s CreatePortalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePortalResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePortalResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePortalResponseBody) GetData() *CreatePortalResponseBodyData {
	return s.Data
}

func (s *CreatePortalResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePortalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePortalResponseBody) SetCode(v string) *CreatePortalResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePortalResponseBody) SetData(v *CreatePortalResponseBodyData) *CreatePortalResponseBody {
	s.Data = v
	return s
}

func (s *CreatePortalResponseBody) SetMessage(v string) *CreatePortalResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePortalResponseBody) SetRequestId(v string) *CreatePortalResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePortalResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePortalResponseBodyData struct {
	// The portal ID. It can be used to locate the portal in API operations such as GetPortal, UpdatePortal, and DeletePortal.
	//
	// example:
	//
	// portal-xxxxxxxx
	PortalId *string `json:"portalId,omitempty" xml:"portalId,omitempty"`
}

func (s CreatePortalResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreatePortalResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePortalResponseBodyData) GetPortalId() *string {
	return s.PortalId
}

func (s *CreatePortalResponseBodyData) SetPortalId(v string) *CreatePortalResponseBodyData {
	s.PortalId = &v
	return s
}

func (s *CreatePortalResponseBodyData) Validate() error {
	return dara.Validate(s)
}
