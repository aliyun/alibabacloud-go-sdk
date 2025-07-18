// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStreamingDataServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorMessage(v string) *DeleteStreamingDataServiceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteStreamingDataServiceResponseBody
	GetRequestId() *string
	SetStatus(v bool) *DeleteStreamingDataServiceResponseBody
	GetStatus() *bool
}

type DeleteStreamingDataServiceResponseBody struct {
	// The error message returned if the operation fails.
	//
	// This parameter is returned only when the return value of **Status*	- is **false**.
	//
	// example:
	//
	// This external service cannot be deleted because it is still used by other data source.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the operation. Valid values:
	//
	// 	- **false**: The operation fails.
	//
	// 	- **true**: The operation is successful.
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteStreamingDataServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteStreamingDataServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStreamingDataServiceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteStreamingDataServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteStreamingDataServiceResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *DeleteStreamingDataServiceResponseBody) SetErrorMessage(v string) *DeleteStreamingDataServiceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteStreamingDataServiceResponseBody) SetRequestId(v string) *DeleteStreamingDataServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStreamingDataServiceResponseBody) SetStatus(v bool) *DeleteStreamingDataServiceResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteStreamingDataServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
