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
	// The code of the face group. If you do not specify this parameter, the face data of the user in all face groups is deleted.
	//
	// example:
	//
	// sgl****7uc
	FaceGroupCode *string `json:"FaceGroupCode,omitempty" xml:"FaceGroupCode,omitempty"`
	// The unique identifier of the user. This value must be the same as the one specified when the face record was added by using AddFaceRecord. If no user identifier was specified when the record was added, you can use the default image name.
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
