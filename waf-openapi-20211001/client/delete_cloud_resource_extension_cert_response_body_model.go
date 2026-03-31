// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudResourceExtensionCertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCloudResourceExtensionCertResponseBody
	GetRequestId() *string
}

type DeleteCloudResourceExtensionCertResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 56B40D30-4960-4F19-B7D5-***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCloudResourceExtensionCertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudResourceExtensionCertResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCloudResourceExtensionCertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCloudResourceExtensionCertResponseBody) SetRequestId(v string) *DeleteCloudResourceExtensionCertResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCloudResourceExtensionCertResponseBody) Validate() error {
	return dara.Validate(s)
}
