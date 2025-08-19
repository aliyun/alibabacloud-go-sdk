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
	// Unique identifier for real-person authentication.
	//
	// example:
	//
	// shae18209d29ce4e8ba252caae98ab15
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// Whether deletion depends on having already obtained relevant data from the corresponding authentication process.
	//
	// - Y: Required. To successfully delete the related data, you must have obtained the processing result through the DescribeFaceVerify interface.
	//
	// - N: Not required (default). For pure server-side API integration, you can directly pass N.
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
