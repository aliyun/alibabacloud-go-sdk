// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTempOssUrlIntlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetObjectName(v string) *TempOssUrlIntlRequest
	GetObjectName() *string
}

type TempOssUrlIntlRequest struct {
	// Object name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250512/pdf/579047.pdf
	ObjectName *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
}

func (s TempOssUrlIntlRequest) String() string {
	return dara.Prettify(s)
}

func (s TempOssUrlIntlRequest) GoString() string {
	return s.String()
}

func (s *TempOssUrlIntlRequest) GetObjectName() *string {
	return s.ObjectName
}

func (s *TempOssUrlIntlRequest) SetObjectName(v string) *TempOssUrlIntlRequest {
	s.ObjectName = &v
	return s
}

func (s *TempOssUrlIntlRequest) Validate() error {
	return dara.Validate(s)
}
