// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOmniSearchResult interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OmniSearchResult
	GetRequestId() *string
	SetSuccess(v string) *OmniSearchResult
	GetSuccess() *string
}

type OmniSearchResult struct {
	// example:
	//
	// 35E5608A-A737-2038-AFB6-D9D34C6BFD9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OmniSearchResult) String() string {
	return dara.Prettify(s)
}

func (s OmniSearchResult) GoString() string {
	return s.String()
}

func (s *OmniSearchResult) GetRequestId() *string {
	return s.RequestId
}

func (s *OmniSearchResult) GetSuccess() *string {
	return s.Success
}

func (s *OmniSearchResult) SetRequestId(v string) *OmniSearchResult {
	s.RequestId = &v
	return s
}

func (s *OmniSearchResult) SetSuccess(v string) *OmniSearchResult {
	s.Success = &v
	return s
}

func (s *OmniSearchResult) Validate() error {
	return dara.Validate(s)
}
