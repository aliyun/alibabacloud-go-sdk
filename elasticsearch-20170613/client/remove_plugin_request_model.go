// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePluginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *RemovePluginRequest
	GetBody() *string
}

type RemovePluginRequest struct {
	// example:
	//
	// [
	//
	//   {
	//
	//     "name": "pluginName",
	//
	//     "elasticsearchVersion": "8.17.0",
	//
	//     "version": "8.17.0",
	//
	//     "fileVersion": "CAEQbxiBgIDMoJqe6hkiIGYzM****",
	//
	//     "state": "UNINSTALLED",
	//
	//     "source": "USER",
	//
	//     "pluginType": "CUSTOM_PLUGIN"
	//
	//   }
	//
	// ]
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemovePluginRequest) String() string {
	return dara.Prettify(s)
}

func (s RemovePluginRequest) GoString() string {
	return s.String()
}

func (s *RemovePluginRequest) GetBody() *string {
	return s.Body
}

func (s *RemovePluginRequest) SetBody(v string) *RemovePluginRequest {
	s.Body = &v
	return s
}

func (s *RemovePluginRequest) Validate() error {
	return dara.Validate(s)
}
