// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudResourceDefaultCertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCloudResourceDefaultCertResponseBody
	GetRequestId() *string
}

type ModifyCloudResourceDefaultCertResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCloudResourceDefaultCertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudResourceDefaultCertResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCloudResourceDefaultCertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCloudResourceDefaultCertResponseBody) SetRequestId(v string) *ModifyCloudResourceDefaultCertResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCloudResourceDefaultCertResponseBody) Validate() error {
	return dara.Validate(s)
}
