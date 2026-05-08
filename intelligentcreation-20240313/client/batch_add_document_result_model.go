// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAddDocumentResult interface {
	dara.Model
	String() string
	GoString() string
	SetAddDocumentResults(v []*AddDocumentResult) *BatchAddDocumentResult
	GetAddDocumentResults() []*AddDocumentResult
	SetRequestId(v string) *BatchAddDocumentResult
	GetRequestId() *string
}

type BatchAddDocumentResult struct {
	// This parameter is required.
	AddDocumentResults []*AddDocumentResult `json:"addDocumentResults,omitempty" xml:"addDocumentResults,omitempty" type:"Repeated"`
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s BatchAddDocumentResult) String() string {
	return dara.Prettify(s)
}

func (s BatchAddDocumentResult) GoString() string {
	return s.String()
}

func (s *BatchAddDocumentResult) GetAddDocumentResults() []*AddDocumentResult {
	return s.AddDocumentResults
}

func (s *BatchAddDocumentResult) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchAddDocumentResult) SetAddDocumentResults(v []*AddDocumentResult) *BatchAddDocumentResult {
	s.AddDocumentResults = v
	return s
}

func (s *BatchAddDocumentResult) SetRequestId(v string) *BatchAddDocumentResult {
	s.RequestId = &v
	return s
}

func (s *BatchAddDocumentResult) Validate() error {
	if s.AddDocumentResults != nil {
		for _, item := range s.AddDocumentResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
