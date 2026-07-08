// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocumentExtractionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *DocumentExtractionShrinkRequest
	GetAgentKey() *string
	SetUrlsShrink(v string) *DocumentExtractionShrinkRequest
	GetUrlsShrink() *string
}

type DocumentExtractionShrinkRequest struct {
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
	UrlsShrink *string `json:"Urls,omitempty" xml:"Urls,omitempty"`
}

func (s DocumentExtractionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DocumentExtractionShrinkRequest) GoString() string {
	return s.String()
}

func (s *DocumentExtractionShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *DocumentExtractionShrinkRequest) GetUrlsShrink() *string {
	return s.UrlsShrink
}

func (s *DocumentExtractionShrinkRequest) SetAgentKey(v string) *DocumentExtractionShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *DocumentExtractionShrinkRequest) SetUrlsShrink(v string) *DocumentExtractionShrinkRequest {
	s.UrlsShrink = &v
	return s
}

func (s *DocumentExtractionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
