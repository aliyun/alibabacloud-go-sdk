// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStreamingDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceId(v int32) *CreateStreamingDataSourceResponseBody
	GetDataSourceId() *int32
	SetRequestId(v string) *CreateStreamingDataSourceResponseBody
	GetRequestId() *string
}

type CreateStreamingDataSourceResponseBody struct {
	// Data source ID.
	//
	// example:
	//
	// 1
	DataSourceId *int32 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// Request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateStreamingDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStreamingDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStreamingDataSourceResponseBody) GetDataSourceId() *int32 {
	return s.DataSourceId
}

func (s *CreateStreamingDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateStreamingDataSourceResponseBody) SetDataSourceId(v int32) *CreateStreamingDataSourceResponseBody {
	s.DataSourceId = &v
	return s
}

func (s *CreateStreamingDataSourceResponseBody) SetRequestId(v string) *CreateStreamingDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStreamingDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
