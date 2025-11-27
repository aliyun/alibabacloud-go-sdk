// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSemanticQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFiles(v []*File) *SemanticQueryResponseBody
	GetFiles() []*File
	SetRequestId(v string) *SemanticQueryResponseBody
	GetRequestId() *string
}

type SemanticQueryResponseBody struct {
	// The files.
	Files []*File `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 2C5C1E0F-D8B8-4DA0-8127-EC32C771****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SemanticQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SemanticQueryResponseBody) GoString() string {
	return s.String()
}

func (s *SemanticQueryResponseBody) GetFiles() []*File {
	return s.Files
}

func (s *SemanticQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SemanticQueryResponseBody) SetFiles(v []*File) *SemanticQueryResponseBody {
	s.Files = v
	return s
}

func (s *SemanticQueryResponseBody) SetRequestId(v string) *SemanticQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *SemanticQueryResponseBody) Validate() error {
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
