// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateCloudResourceResponseBody
	GetRequestId() *string
}

type CreateCloudResourceResponseBody struct {
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateCloudResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCloudResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCloudResourceResponseBody) SetRequestId(v string) *CreateCloudResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCloudResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
