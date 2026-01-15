// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareSimilarByImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *CompareSimilarByImageRequest
	GetInstanceName() *string
	SetPrimaryPicContent(v string) *CompareSimilarByImageRequest
	GetPrimaryPicContent() *string
	SetSecondaryPicContent(v string) *CompareSimilarByImageRequest
	GetSecondaryPicContent() *string
}

type CompareSimilarByImageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// demoinstance1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AAAANSUhEUgAAAPcAAAEVCAYAAAA8d3NuAAAAAXNSR0IArs......RK5CYII=
	PrimaryPicContent *string `json:"PrimaryPicContent,omitempty" xml:"PrimaryPicContent,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AAAANSUhEUgAAAPcAAAEVCAYAAAA8d3NuAAAAAXNSR0IArs......RK5CYII=
	SecondaryPicContent *string `json:"SecondaryPicContent,omitempty" xml:"SecondaryPicContent,omitempty"`
}

func (s CompareSimilarByImageRequest) String() string {
	return dara.Prettify(s)
}

func (s CompareSimilarByImageRequest) GoString() string {
	return s.String()
}

func (s *CompareSimilarByImageRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CompareSimilarByImageRequest) GetPrimaryPicContent() *string {
	return s.PrimaryPicContent
}

func (s *CompareSimilarByImageRequest) GetSecondaryPicContent() *string {
	return s.SecondaryPicContent
}

func (s *CompareSimilarByImageRequest) SetInstanceName(v string) *CompareSimilarByImageRequest {
	s.InstanceName = &v
	return s
}

func (s *CompareSimilarByImageRequest) SetPrimaryPicContent(v string) *CompareSimilarByImageRequest {
	s.PrimaryPicContent = &v
	return s
}

func (s *CompareSimilarByImageRequest) SetSecondaryPicContent(v string) *CompareSimilarByImageRequest {
	s.SecondaryPicContent = &v
	return s
}

func (s *CompareSimilarByImageRequest) Validate() error {
	return dara.Validate(s)
}
