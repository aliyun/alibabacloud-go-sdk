// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudResourceExtensionCertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateCloudResourceExtensionCertResponseBody
	GetRequestId() *string
}

type CreateCloudResourceExtensionCertResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19****5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCloudResourceExtensionCertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudResourceExtensionCertResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCloudResourceExtensionCertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCloudResourceExtensionCertResponseBody) SetRequestId(v string) *CreateCloudResourceExtensionCertResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCloudResourceExtensionCertResponseBody) Validate() error {
	return dara.Validate(s)
}
