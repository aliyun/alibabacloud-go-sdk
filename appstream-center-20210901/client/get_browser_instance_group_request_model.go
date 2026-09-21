// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBrowserInstanceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrowserInstanceGroupId(v string) *GetBrowserInstanceGroupRequest
	GetBrowserInstanceGroupId() *string
}

type GetBrowserInstanceGroupRequest struct {
	// The cloud browser group ID. This parameter is required. Specify the ID of a browser group that is created under the current account.
	//
	// example:
	//
	// big-0c7loey7fzjq****
	BrowserInstanceGroupId *string `json:"BrowserInstanceGroupId,omitempty" xml:"BrowserInstanceGroupId,omitempty"`
}

func (s GetBrowserInstanceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBrowserInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *GetBrowserInstanceGroupRequest) GetBrowserInstanceGroupId() *string {
	return s.BrowserInstanceGroupId
}

func (s *GetBrowserInstanceGroupRequest) SetBrowserInstanceGroupId(v string) *GetBrowserInstanceGroupRequest {
	s.BrowserInstanceGroupId = &v
	return s
}

func (s *GetBrowserInstanceGroupRequest) Validate() error {
	return dara.Validate(s)
}
