// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverRenderingDataPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RecoverRenderingDataPackageResponseBody
	GetRequestId() *string
}

type RecoverRenderingDataPackageResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecoverRenderingDataPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecoverRenderingDataPackageResponseBody) GoString() string {
	return s.String()
}

func (s *RecoverRenderingDataPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecoverRenderingDataPackageResponseBody) SetRequestId(v string) *RecoverRenderingDataPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecoverRenderingDataPackageResponseBody) Validate() error {
	return dara.Validate(s)
}
