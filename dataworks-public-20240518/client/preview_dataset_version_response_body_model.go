// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewDatasetVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPreviewResult(v *PreviewDatasetVersionResponseBodyPreviewResult) *PreviewDatasetVersionResponseBody
	GetPreviewResult() *PreviewDatasetVersionResponseBodyPreviewResult
	SetRequestId(v string) *PreviewDatasetVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PreviewDatasetVersionResponseBody
	GetSuccess() *bool
}

type PreviewDatasetVersionResponseBody struct {
	// Preview results
	PreviewResult *PreviewDatasetVersionResponseBodyPreviewResult `json:"PreviewResult,omitempty" xml:"PreviewResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// A6C6B486-E3A2-5D52-9E76-D9380485DXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PreviewDatasetVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PreviewDatasetVersionResponseBody) GoString() string {
	return s.String()
}

func (s *PreviewDatasetVersionResponseBody) GetPreviewResult() *PreviewDatasetVersionResponseBodyPreviewResult {
	return s.PreviewResult
}

func (s *PreviewDatasetVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PreviewDatasetVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PreviewDatasetVersionResponseBody) SetPreviewResult(v *PreviewDatasetVersionResponseBodyPreviewResult) *PreviewDatasetVersionResponseBody {
	s.PreviewResult = v
	return s
}

func (s *PreviewDatasetVersionResponseBody) SetRequestId(v string) *PreviewDatasetVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreviewDatasetVersionResponseBody) SetSuccess(v bool) *PreviewDatasetVersionResponseBody {
	s.Success = &v
	return s
}

func (s *PreviewDatasetVersionResponseBody) Validate() error {
	if s.PreviewResult != nil {
		if err := s.PreviewResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PreviewDatasetVersionResponseBodyPreviewResult struct {
	// Content
	//
	// example:
	//
	// this is content
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// File name
	//
	// example:
	//
	// parth/data.csv
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The MIME type
	//
	// example:
	//
	// text/plain
	MimeType *string `json:"MimeType,omitempty" xml:"MimeType,omitempty"`
	// Preview availability
	//
	// example:
	//
	// true
	SupportPreview *bool `json:"SupportPreview,omitempty" xml:"SupportPreview,omitempty"`
}

func (s PreviewDatasetVersionResponseBodyPreviewResult) String() string {
	return dara.Prettify(s)
}

func (s PreviewDatasetVersionResponseBodyPreviewResult) GoString() string {
	return s.String()
}

func (s *PreviewDatasetVersionResponseBodyPreviewResult) GetContent() *string {
	return s.Content
}

func (s *PreviewDatasetVersionResponseBodyPreviewResult) GetFileName() *string {
	return s.FileName
}

func (s *PreviewDatasetVersionResponseBodyPreviewResult) GetMimeType() *string {
	return s.MimeType
}

func (s *PreviewDatasetVersionResponseBodyPreviewResult) GetSupportPreview() *bool {
	return s.SupportPreview
}

func (s *PreviewDatasetVersionResponseBodyPreviewResult) SetContent(v string) *PreviewDatasetVersionResponseBodyPreviewResult {
	s.Content = &v
	return s
}

func (s *PreviewDatasetVersionResponseBodyPreviewResult) SetFileName(v string) *PreviewDatasetVersionResponseBodyPreviewResult {
	s.FileName = &v
	return s
}

func (s *PreviewDatasetVersionResponseBodyPreviewResult) SetMimeType(v string) *PreviewDatasetVersionResponseBodyPreviewResult {
	s.MimeType = &v
	return s
}

func (s *PreviewDatasetVersionResponseBodyPreviewResult) SetSupportPreview(v bool) *PreviewDatasetVersionResponseBodyPreviewResult {
	s.SupportPreview = &v
	return s
}

func (s *PreviewDatasetVersionResponseBodyPreviewResult) Validate() error {
	return dara.Validate(s)
}
