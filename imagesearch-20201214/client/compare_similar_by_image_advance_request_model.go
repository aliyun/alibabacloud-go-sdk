// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iCompareSimilarByImageAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *CompareSimilarByImageAdvanceRequest
	GetInstanceName() *string
	SetPrimaryPicContentObject(v io.Reader) *CompareSimilarByImageAdvanceRequest
	GetPrimaryPicContentObject() io.Reader
	SetSecondaryPicContentObject(v io.Reader) *CompareSimilarByImageAdvanceRequest
	GetSecondaryPicContentObject() io.Reader
}

type CompareSimilarByImageAdvanceRequest struct {
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
	PrimaryPicContentObject io.Reader `json:"PrimaryPicContent,omitempty" xml:"PrimaryPicContent,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AAAANSUhEUgAAAPcAAAEVCAYAAAA8d3NuAAAAAXNSR0IArs......RK5CYII=
	SecondaryPicContentObject io.Reader `json:"SecondaryPicContent,omitempty" xml:"SecondaryPicContent,omitempty"`
}

func (s CompareSimilarByImageAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CompareSimilarByImageAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CompareSimilarByImageAdvanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CompareSimilarByImageAdvanceRequest) GetPrimaryPicContentObject() io.Reader {
	return s.PrimaryPicContentObject
}

func (s *CompareSimilarByImageAdvanceRequest) GetSecondaryPicContentObject() io.Reader {
	return s.SecondaryPicContentObject
}

func (s *CompareSimilarByImageAdvanceRequest) SetInstanceName(v string) *CompareSimilarByImageAdvanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CompareSimilarByImageAdvanceRequest) SetPrimaryPicContentObject(v io.Reader) *CompareSimilarByImageAdvanceRequest {
	s.PrimaryPicContentObject = v
	return s
}

func (s *CompareSimilarByImageAdvanceRequest) SetSecondaryPicContentObject(v io.Reader) *CompareSimilarByImageAdvanceRequest {
	s.SecondaryPicContentObject = v
	return s
}

func (s *CompareSimilarByImageAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
