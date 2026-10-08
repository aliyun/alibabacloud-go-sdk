// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMosCheckInRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityId(v string) *MosCheckInRequest
	GetActivityId() *string
	SetExtParam(v string) *MosCheckInRequest
	GetExtParam() *string
	SetQrCode(v string) *MosCheckInRequest
	GetQrCode() *string
}

type MosCheckInRequest struct {
	// example:
	//
	// INTL1234
	ActivityId *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// example:
	//
	// {}
	ExtParam *string `json:"ExtParam,omitempty" xml:"ExtParam,omitempty"`
	// example:
	//
	// abc12345
	QrCode *string `json:"QrCode,omitempty" xml:"QrCode,omitempty"`
}

func (s MosCheckInRequest) String() string {
	return dara.Prettify(s)
}

func (s MosCheckInRequest) GoString() string {
	return s.String()
}

func (s *MosCheckInRequest) GetActivityId() *string {
	return s.ActivityId
}

func (s *MosCheckInRequest) GetExtParam() *string {
	return s.ExtParam
}

func (s *MosCheckInRequest) GetQrCode() *string {
	return s.QrCode
}

func (s *MosCheckInRequest) SetActivityId(v string) *MosCheckInRequest {
	s.ActivityId = &v
	return s
}

func (s *MosCheckInRequest) SetExtParam(v string) *MosCheckInRequest {
	s.ExtParam = &v
	return s
}

func (s *MosCheckInRequest) SetQrCode(v string) *MosCheckInRequest {
	s.QrCode = &v
	return s
}

func (s *MosCheckInRequest) Validate() error {
	return dara.Validate(s)
}
