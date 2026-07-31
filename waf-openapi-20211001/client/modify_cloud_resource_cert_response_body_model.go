// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudResourceCertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCloudResourceCertResponseBody
	GetRequestId() *string
}

type ModifyCloudResourceCertResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// A47D2DC0-7151-51EF-BFB7-***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCloudResourceCertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudResourceCertResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCloudResourceCertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCloudResourceCertResponseBody) SetRequestId(v string) *ModifyCloudResourceCertResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCloudResourceCertResponseBody) Validate() error {
	return dara.Validate(s)
}
