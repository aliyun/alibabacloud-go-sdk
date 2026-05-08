// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocumentResult interface {
	dara.Model
	String() string
	GoString() string
	SetDocumentInfo(v *DocumentInfo) *DocumentResult
	GetDocumentInfo() *DocumentInfo
	SetRequestId(v string) *DocumentResult
	GetRequestId() *string
}

type DocumentResult struct {
	DocumentInfo *DocumentInfo `json:"documentInfo,omitempty" xml:"documentInfo,omitempty"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DocumentResult) String() string {
	return dara.Prettify(s)
}

func (s DocumentResult) GoString() string {
	return s.String()
}

func (s *DocumentResult) GetDocumentInfo() *DocumentInfo {
	return s.DocumentInfo
}

func (s *DocumentResult) GetRequestId() *string {
	return s.RequestId
}

func (s *DocumentResult) SetDocumentInfo(v *DocumentInfo) *DocumentResult {
	s.DocumentInfo = v
	return s
}

func (s *DocumentResult) SetRequestId(v string) *DocumentResult {
	s.RequestId = &v
	return s
}

func (s *DocumentResult) Validate() error {
	if s.DocumentInfo != nil {
		if err := s.DocumentInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}
