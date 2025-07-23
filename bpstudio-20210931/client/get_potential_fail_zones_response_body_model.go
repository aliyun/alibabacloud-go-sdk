// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPotentialFailZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPotentialFailZonesResponseBody
	GetCode() *string
	SetData(v []*string) *GetPotentialFailZonesResponseBody
	GetData() []*string
	SetMessage(v string) *GetPotentialFailZonesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPotentialFailZonesResponseBody
	GetRequestId() *string
}

type GetPotentialFailZonesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The zones where the current disaster recovery service can be switched.
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error message.
	//
	// example:
	//
	// The specified ResourceIds are not found in our records.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BFB7F5C8-FE7A-06CA-9F87-ABBF6B848F0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPotentialFailZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPotentialFailZonesResponseBody) GoString() string {
	return s.String()
}

func (s *GetPotentialFailZonesResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPotentialFailZonesResponseBody) GetData() []*string {
	return s.Data
}

func (s *GetPotentialFailZonesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPotentialFailZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPotentialFailZonesResponseBody) SetCode(v string) *GetPotentialFailZonesResponseBody {
	s.Code = &v
	return s
}

func (s *GetPotentialFailZonesResponseBody) SetData(v []*string) *GetPotentialFailZonesResponseBody {
	s.Data = v
	return s
}

func (s *GetPotentialFailZonesResponseBody) SetMessage(v string) *GetPotentialFailZonesResponseBody {
	s.Message = &v
	return s
}

func (s *GetPotentialFailZonesResponseBody) SetRequestId(v string) *GetPotentialFailZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPotentialFailZonesResponseBody) Validate() error {
	return dara.Validate(s)
}
