// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iGenerateDynamicImageAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperation(v string) *GenerateDynamicImageAdvanceRequest
	GetOperation() *string
	SetUrlObject(v io.Reader) *GenerateDynamicImageAdvanceRequest
	GetUrlObject() io.Reader
}

type GenerateDynamicImageAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 11
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://vigen-invi.oss-cn-shanghai.aliyuncs.com/temp/xl/dynamicPhoto/viapi_test_images/xxxx.jpg
	UrlObject io.Reader `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GenerateDynamicImageAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateDynamicImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *GenerateDynamicImageAdvanceRequest) GetOperation() *string {
	return s.Operation
}

func (s *GenerateDynamicImageAdvanceRequest) GetUrlObject() io.Reader {
	return s.UrlObject
}

func (s *GenerateDynamicImageAdvanceRequest) SetOperation(v string) *GenerateDynamicImageAdvanceRequest {
	s.Operation = &v
	return s
}

func (s *GenerateDynamicImageAdvanceRequest) SetUrlObject(v io.Reader) *GenerateDynamicImageAdvanceRequest {
	s.UrlObject = v
	return s
}

func (s *GenerateDynamicImageAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
