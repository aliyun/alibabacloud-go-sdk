// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSuggestShrinkableNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSuggestShrinkableNodesResponseBody
	GetRequestId() *string
	SetResult(v []*GetSuggestShrinkableNodesResponseBodyResult) *GetSuggestShrinkableNodesResponseBody
	GetResult() []*GetSuggestShrinkableNodesResponseBodyResult
}

type GetSuggestShrinkableNodesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The return results.
	Result []*GetSuggestShrinkableNodesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s GetSuggestShrinkableNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSuggestShrinkableNodesResponseBody) GoString() string {
	return s.String()
}

func (s *GetSuggestShrinkableNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSuggestShrinkableNodesResponseBody) GetResult() []*GetSuggestShrinkableNodesResponseBodyResult {
	return s.Result
}

func (s *GetSuggestShrinkableNodesResponseBody) SetRequestId(v string) *GetSuggestShrinkableNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSuggestShrinkableNodesResponseBody) SetResult(v []*GetSuggestShrinkableNodesResponseBodyResult) *GetSuggestShrinkableNodesResponseBody {
	s.Result = v
	return s
}

func (s *GetSuggestShrinkableNodesResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSuggestShrinkableNodesResponseBodyResult struct {
	// The IP address of the node.
	//
	// example:
	//
	// ``192.168.**.**``
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// The access port number of the node.
	//
	// example:
	//
	// 9200
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
}

func (s GetSuggestShrinkableNodesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetSuggestShrinkableNodesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetSuggestShrinkableNodesResponseBodyResult) GetHost() *string {
	return s.Host
}

func (s *GetSuggestShrinkableNodesResponseBodyResult) GetPort() *int32 {
	return s.Port
}

func (s *GetSuggestShrinkableNodesResponseBodyResult) SetHost(v string) *GetSuggestShrinkableNodesResponseBodyResult {
	s.Host = &v
	return s
}

func (s *GetSuggestShrinkableNodesResponseBodyResult) SetPort(v int32) *GetSuggestShrinkableNodesResponseBodyResult {
	s.Port = &v
	return s
}

func (s *GetSuggestShrinkableNodesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
