// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iDetectKitchenAnimalsAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLAObject(v io.Reader) *DetectKitchenAnimalsAdvanceRequest
	GetImageURLAObject() io.Reader
	SetImageURLBObject(v io.Reader) *DetectKitchenAnimalsAdvanceRequest
	GetImageURLBObject() io.Reader
}

type DetectKitchenAnimalsAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/objectdet/DetectKitchenAnimals/DetectKitchenAnimals-left1.png
	ImageURLAObject io.Reader `json:"ImageURLA,omitempty" xml:"ImageURLA,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/objectdet/DetectKitchenAnimals/DetectKitchenAnimals-right1.png
	ImageURLBObject io.Reader `json:"ImageURLB,omitempty" xml:"ImageURLB,omitempty"`
}

func (s DetectKitchenAnimalsAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectKitchenAnimalsAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DetectKitchenAnimalsAdvanceRequest) GetImageURLAObject() io.Reader {
	return s.ImageURLAObject
}

func (s *DetectKitchenAnimalsAdvanceRequest) GetImageURLBObject() io.Reader {
	return s.ImageURLBObject
}

func (s *DetectKitchenAnimalsAdvanceRequest) SetImageURLAObject(v io.Reader) *DetectKitchenAnimalsAdvanceRequest {
	s.ImageURLAObject = v
	return s
}

func (s *DetectKitchenAnimalsAdvanceRequest) SetImageURLBObject(v io.Reader) *DetectKitchenAnimalsAdvanceRequest {
	s.ImageURLBObject = v
	return s
}

func (s *DetectKitchenAnimalsAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
