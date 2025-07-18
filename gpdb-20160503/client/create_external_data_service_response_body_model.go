// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExternalDataServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateExternalDataServiceResponseBody
	GetRequestId() *string
	SetServiceId(v int32) *CreateExternalDataServiceResponseBody
	GetServiceId() *int32
}

type CreateExternalDataServiceResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Service ID.
	//
	// example:
	//
	// 100
	ServiceId *int32 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s CreateExternalDataServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalDataServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExternalDataServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateExternalDataServiceResponseBody) GetServiceId() *int32 {
	return s.ServiceId
}

func (s *CreateExternalDataServiceResponseBody) SetRequestId(v string) *CreateExternalDataServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExternalDataServiceResponseBody) SetServiceId(v int32) *CreateExternalDataServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *CreateExternalDataServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
