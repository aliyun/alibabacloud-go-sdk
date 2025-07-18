// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStreamingDataServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyStreamingDataServiceResponseBody
	GetRequestId() *string
}

type ModifyStreamingDataServiceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyStreamingDataServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyStreamingDataServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyStreamingDataServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyStreamingDataServiceResponseBody) SetRequestId(v string) *ModifyStreamingDataServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyStreamingDataServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
