// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocumentExtractionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DocumentExtractionRequest
	GetAgentKey() *string
	SetUrls(v []*string) *DocumentExtractionRequest
	GetUrls() []*string
}

type DocumentExtractionRequest struct {
	// Unique identifier of the workspace: [AgentKey](https://help.aliyun.com/document_detail/2587494.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// List of URLs to extract.
	//
	// This parameter is required.
	Urls []*string `json:"Urls,omitempty" xml:"Urls,omitempty" type:"Repeated"`
}

func (s DocumentExtractionRequest) String() string {
	return dara.Prettify(s)
}

func (s DocumentExtractionRequest) GoString() string {
	return s.String()
}

func (s *DocumentExtractionRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DocumentExtractionRequest) GetUrls() []*string {
	return s.Urls
}

func (s *DocumentExtractionRequest) SetAgentKey(v string) *DocumentExtractionRequest {
	s.AgentKey = &v
	return s
}

func (s *DocumentExtractionRequest) SetUrls(v []*string) *DocumentExtractionRequest {
	s.Urls = v
	return s
}

func (s *DocumentExtractionRequest) Validate() error {
	return dara.Validate(s)
}
