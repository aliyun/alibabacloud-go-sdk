// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHadoopDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteHadoopDataSourceResponseBody
	GetRequestId() *string
}

type DeleteHadoopDataSourceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHadoopDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHadoopDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHadoopDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHadoopDataSourceResponseBody) SetRequestId(v string) *DeleteHadoopDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHadoopDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
