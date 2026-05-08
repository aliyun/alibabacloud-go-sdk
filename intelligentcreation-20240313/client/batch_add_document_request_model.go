// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAddDocumentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddDocumentInfos(v []*AddDocumentInfo) *BatchAddDocumentRequest
	GetAddDocumentInfos() []*AddDocumentInfo
}

type BatchAddDocumentRequest struct {
	AddDocumentInfos []*AddDocumentInfo `json:"addDocumentInfos,omitempty" xml:"addDocumentInfos,omitempty" type:"Repeated"`
}

func (s BatchAddDocumentRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchAddDocumentRequest) GoString() string {
	return s.String()
}

func (s *BatchAddDocumentRequest) GetAddDocumentInfos() []*AddDocumentInfo {
	return s.AddDocumentInfos
}

func (s *BatchAddDocumentRequest) SetAddDocumentInfos(v []*AddDocumentInfo) *BatchAddDocumentRequest {
	s.AddDocumentInfos = v
	return s
}

func (s *BatchAddDocumentRequest) Validate() error {
	if s.AddDocumentInfos != nil {
		for _, item := range s.AddDocumentInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
