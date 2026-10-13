// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePortalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdatePortalResponseBody
	GetCode() *string
	SetData(v *UpdatePortalResponseBodyData) *UpdatePortalResponseBody
	GetData() *UpdatePortalResponseBodyData
	SetMessage(v string) *UpdatePortalResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdatePortalResponseBody
	GetRequestId() *string
}

type UpdatePortalResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *UpdatePortalResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s UpdatePortalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePortalResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePortalResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdatePortalResponseBody) GetData() *UpdatePortalResponseBodyData {
	return s.Data
}

func (s *UpdatePortalResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdatePortalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePortalResponseBody) SetCode(v string) *UpdatePortalResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePortalResponseBody) SetData(v *UpdatePortalResponseBodyData) *UpdatePortalResponseBody {
	s.Data = v
	return s
}

func (s *UpdatePortalResponseBody) SetMessage(v string) *UpdatePortalResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePortalResponseBody) SetRequestId(v string) *UpdatePortalResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePortalResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePortalResponseBodyData struct {
	// The ID of the updated portal. This is the primary ID of the portal resource specified by the portalId path parameter.
	//
	// example:
	//
	// portal-xxxxxxxx
	PortalId *string `json:"portalId,omitempty" xml:"portalId,omitempty"`
}

func (s UpdatePortalResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdatePortalResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdatePortalResponseBodyData) GetPortalId() *string {
	return s.PortalId
}

func (s *UpdatePortalResponseBodyData) SetPortalId(v string) *UpdatePortalResponseBodyData {
	s.PortalId = &v
	return s
}

func (s *UpdatePortalResponseBodyData) Validate() error {
	return dara.Validate(s)
}
