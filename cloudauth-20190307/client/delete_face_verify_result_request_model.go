// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceVerifyResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertifyId(v string) *DeleteFaceVerifyResultRequest
	GetCertifyId() *string
	SetDeleteAfterQuery(v string) *DeleteFaceVerifyResultRequest
	GetDeleteAfterQuery() *string
}

type DeleteFaceVerifyResultRequest struct {
	// The unique identifier for ID Verification.
	//
	// example:
	//
	// shae18209d29ce4e8ba252caae98ab15
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// Specifies whether the deletion relies on the data retrieved from the corresponding verification process. Valid values:
	//
	// - Y: Yes. To successfully delete the related data, you must have obtained the processing result by calling the DescribeFaceVerify operation.
	//
	// - N: No. This is the default value. You can directly set this parameter to N when you use the server-only API integration mode.
	//
	// example:
	//
	// Y
	DeleteAfterQuery *string `json:"DeleteAfterQuery,omitempty" xml:"DeleteAfterQuery,omitempty"`
}

func (s DeleteFaceVerifyResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceVerifyResultRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceVerifyResultRequest) GetCertifyId() *string {
	return s.CertifyId
}

func (s *DeleteFaceVerifyResultRequest) GetDeleteAfterQuery() *string {
	return s.DeleteAfterQuery
}

func (s *DeleteFaceVerifyResultRequest) SetCertifyId(v string) *DeleteFaceVerifyResultRequest {
	s.CertifyId = &v
	return s
}

func (s *DeleteFaceVerifyResultRequest) SetDeleteAfterQuery(v string) *DeleteFaceVerifyResultRequest {
	s.DeleteAfterQuery = &v
	return s
}

func (s *DeleteFaceVerifyResultRequest) Validate() error {
	return dara.Validate(s)
}
