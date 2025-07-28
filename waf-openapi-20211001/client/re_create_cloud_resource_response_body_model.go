// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReCreateCloudResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReCreateCloudResourceResponseBody
	GetRequestId() *string
}

type ReCreateCloudResourceResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19****5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReCreateCloudResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReCreateCloudResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ReCreateCloudResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReCreateCloudResourceResponseBody) SetRequestId(v string) *ReCreateCloudResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReCreateCloudResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
