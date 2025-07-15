// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPluginId(v string) *CreatePluginResponseBody
	GetPluginId() *string
	SetRequestId(v string) *CreatePluginResponseBody
	GetRequestId() *string
	SetTagStatus(v bool) *CreatePluginResponseBody
	GetTagStatus() *bool
}

type CreatePluginResponseBody struct {
	// The ID of the plug-in.
	//
	// example:
	//
	// 1f3bde29b43d4d53989248327ff737f2
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// EF924FE4-2EDD-4CD3-89EC-34E4708574E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the plug-in is successfully marked.
	//
	// example:
	//
	// true
	TagStatus *bool `json:"TagStatus,omitempty" xml:"TagStatus,omitempty"`
}

func (s CreatePluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePluginResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePluginResponseBody) GetPluginId() *string {
	return s.PluginId
}

func (s *CreatePluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePluginResponseBody) GetTagStatus() *bool {
	return s.TagStatus
}

func (s *CreatePluginResponseBody) SetPluginId(v string) *CreatePluginResponseBody {
	s.PluginId = &v
	return s
}

func (s *CreatePluginResponseBody) SetRequestId(v string) *CreatePluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePluginResponseBody) SetTagStatus(v bool) *CreatePluginResponseBody {
	s.TagStatus = &v
	return s
}

func (s *CreatePluginResponseBody) Validate() error {
	return dara.Validate(s)
}
