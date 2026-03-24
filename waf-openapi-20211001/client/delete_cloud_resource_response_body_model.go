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
	// The ID of the request. Alibaba Cloud generates a unique ID for each request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
