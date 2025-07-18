// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStreamingDataServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateStreamingDataServiceResponseBody
	GetRequestId() *string
	SetServiceId(v int32) *CreateStreamingDataServiceResponseBody
	GetServiceId() *int32
}

type CreateStreamingDataServiceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The real-time data service ID.
	//
	// example:
	//
	// 1
	ServiceId *int32 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s CreateStreamingDataServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStreamingDataServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStreamingDataServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateStreamingDataServiceResponseBody) GetServiceId() *int32 {
	return s.ServiceId
}

func (s *CreateStreamingDataServiceResponseBody) SetRequestId(v string) *CreateStreamingDataServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStreamingDataServiceResponseBody) SetServiceId(v int32) *CreateStreamingDataServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *CreateStreamingDataServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
