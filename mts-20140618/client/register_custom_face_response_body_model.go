// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterCustomFaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFaceId(v string) *RegisterCustomFaceResponseBody
	GetFaceId() *string
	SetRequestId(v string) *RegisterCustomFaceResponseBody
	GetRequestId() *string
}

type RegisterCustomFaceResponseBody struct {
	// The ID of the face.
	//
	// example:
	//
	// c6cc71cb44a9491093818faf9d60****
	FaceId *string `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 91AEA76D-25B5-50DF-9126-AA6BB10FDAF4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterCustomFaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterCustomFaceResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterCustomFaceResponseBody) GetFaceId() *string {
	return s.FaceId
}

func (s *RegisterCustomFaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterCustomFaceResponseBody) SetFaceId(v string) *RegisterCustomFaceResponseBody {
	s.FaceId = &v
	return s
}

func (s *RegisterCustomFaceResponseBody) SetRequestId(v string) *RegisterCustomFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterCustomFaceResponseBody) Validate() error {
	return dara.Validate(s)
}
