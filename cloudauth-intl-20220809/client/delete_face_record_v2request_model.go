// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceRecordV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetFaceGroupCode(v string) *DeleteFaceRecordV2Request
	GetFaceGroupCode() *string
	SetMerchantUserId(v string) *DeleteFaceRecordV2Request
	GetMerchantUserId() *string
}

type DeleteFaceRecordV2Request struct {
	// The face group code. If this parameter is not specified, the face data of the user is deleted from all face groups.
	//
	// example:
	//
	// sgl****7uc
	FaceGroupCode *string `json:"FaceGroupCode,omitempty" xml:"FaceGroupCode,omitempty"`
	// The unique user identifier, which must be consistent with the one used when calling AddFaceRecord. If this parameter was not specified during registration, you can use the default image name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1231****
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
}

func (s DeleteFaceRecordV2Request) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceRecordV2Request) GoString() string {
	return s.String()
}

func (s *DeleteFaceRecordV2Request) GetFaceGroupCode() *string {
	return s.FaceGroupCode
}

func (s *DeleteFaceRecordV2Request) GetMerchantUserId() *string {
	return s.MerchantUserId
}

func (s *DeleteFaceRecordV2Request) SetFaceGroupCode(v string) *DeleteFaceRecordV2Request {
	s.FaceGroupCode = &v
	return s
}

func (s *DeleteFaceRecordV2Request) SetMerchantUserId(v string) *DeleteFaceRecordV2Request {
	s.MerchantUserId = &v
	return s
}

func (s *DeleteFaceRecordV2Request) Validate() error {
	return dara.Validate(s)
}
