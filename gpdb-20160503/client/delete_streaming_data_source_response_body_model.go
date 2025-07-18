// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStreamingDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteStreamingDataSourceResponseBody
	GetRequestId() *string
}

type DeleteStreamingDataSourceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteStreamingDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteStreamingDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStreamingDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteStreamingDataSourceResponseBody) SetRequestId(v string) *DeleteStreamingDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStreamingDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
