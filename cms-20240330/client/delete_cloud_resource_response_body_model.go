// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCloudResourceResponseBody
	GetRequestId() *string
}

type DeleteCloudResourceResponseBody struct {
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteCloudResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCloudResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCloudResourceResponseBody) SetRequestId(v string) *DeleteCloudResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCloudResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
