// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRunLabelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRunLabelResponseBody
	GetRequestId() *string
}

type DeleteRunLabelResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ADF6D849-*****-7E7030F0CE53
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteRunLabelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRunLabelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRunLabelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRunLabelResponseBody) SetRequestId(v string) *DeleteRunLabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRunLabelResponseBody) Validate() error {
	return dara.Validate(s)
}
