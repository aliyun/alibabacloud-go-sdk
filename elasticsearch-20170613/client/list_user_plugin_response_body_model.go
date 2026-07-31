// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserPluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]interface{}) *ListUserPluginResponseBody
	GetHeaders() map[string]interface{}
	SetRequestId(v string) *ListUserPluginResponseBody
	GetRequestId() *string
	SetResult(v []*ListUserPluginResponseBodyResult) *ListUserPluginResponseBody
	GetResult() []*ListUserPluginResponseBodyResult
}

type ListUserPluginResponseBody struct {
	// The response headers.
	//
	// example:
	//
	// {\\"totalCount\\": 1, \\"X-Total-Count\\": 1}
	Headers map[string]interface{} `json:"Headers,omitempty" xml:"Headers,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5EEF8FAE-EEDD***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result array.
	Result []*ListUserPluginResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListUserPluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserPluginResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserPluginResponseBody) GetHeaders() map[string]interface{} {
	return s.Headers
}

func (s *ListUserPluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserPluginResponseBody) GetResult() []*ListUserPluginResponseBodyResult {
	return s.Result
}

func (s *ListUserPluginResponseBody) SetHeaders(v map[string]interface{}) *ListUserPluginResponseBody {
	s.Headers = v
	return s
}

func (s *ListUserPluginResponseBody) SetRequestId(v string) *ListUserPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserPluginResponseBody) SetResult(v []*ListUserPluginResponseBodyResult) *ListUserPluginResponseBody {
	s.Result = v
	return s
}

func (s *ListUserPluginResponseBody) Validate() error {
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

type ListUserPluginResponseBodyResult struct {
	// The list of plug-ins with the same name.
	BingoPlugins []*ListUserPluginResponseBodyResultBingoPlugins `json:"bingoPlugins,omitempty" xml:"bingoPlugins,omitempty" type:"Repeated"`
	// The plug-in name.
	//
	// example:
	//
	// ct-test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The plug-in source.
	//
	// example:
	//
	// USER
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The plug-in status.
	//
	// example:
	//
	// UNINSTALLED
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The plug-in version.
	//
	// example:
	//
	// 8.17.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListUserPluginResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListUserPluginResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListUserPluginResponseBodyResult) GetBingoPlugins() []*ListUserPluginResponseBodyResultBingoPlugins {
	return s.BingoPlugins
}

func (s *ListUserPluginResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListUserPluginResponseBodyResult) GetSource() *string {
	return s.Source
}

func (s *ListUserPluginResponseBodyResult) GetState() *string {
	return s.State
}

func (s *ListUserPluginResponseBodyResult) GetVersion() *string {
	return s.Version
}

func (s *ListUserPluginResponseBodyResult) SetBingoPlugins(v []*ListUserPluginResponseBodyResultBingoPlugins) *ListUserPluginResponseBodyResult {
	s.BingoPlugins = v
	return s
}

func (s *ListUserPluginResponseBodyResult) SetName(v string) *ListUserPluginResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListUserPluginResponseBodyResult) SetSource(v string) *ListUserPluginResponseBodyResult {
	s.Source = &v
	return s
}

func (s *ListUserPluginResponseBodyResult) SetState(v string) *ListUserPluginResponseBodyResult {
	s.State = &v
	return s
}

func (s *ListUserPluginResponseBodyResult) SetVersion(v string) *ListUserPluginResponseBodyResult {
	s.Version = &v
	return s
}

func (s *ListUserPluginResponseBodyResult) Validate() error {
	if s.BingoPlugins != nil {
		for _, item := range s.BingoPlugins {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserPluginResponseBodyResultBingoPlugins struct {
	// The plug-in description.
	//
	// example:
	//
	// The plugin***
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The Elasticsearch version of the plug-in.
	//
	// example:
	//
	// 8.17.0
	ElasticsearchVersion *string `json:"elasticsearchVersion,omitempty" xml:"elasticsearchVersion,omitempty"`
	// The unique identifier of the plug-in.
	//
	// example:
	//
	// CAEQaRiBgIDI2tie6hkiIGIwM2I3MjZmNjk3YzR***
	FileVersion *string `json:"fileVersion,omitempty" xml:"fileVersion,omitempty"`
	// The plug-in name.
	//
	// example:
	//
	// dynamic-name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The plug-in source. Valid values:
	//
	// - USER: custom plug-in.
	//
	// - SYSTEM: system preset plug-in.
	//
	// example:
	//
	// USER
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The plug-in status. Valid values:
	//
	// - INSTALLED
	//
	// - UNINSTALLED
	//
	// - INSTALLING
	//
	// - UNINSTALLING
	//
	// - UPGRADING
	//
	// - FAILED
	//
	// - UNKNOWN
	//
	// - UPLOADING
	//
	// example:
	//
	// UNINSTALLED
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The plug-in version.
	//
	// example:
	//
	// 8.17.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListUserPluginResponseBodyResultBingoPlugins) String() string {
	return dara.Prettify(s)
}

func (s ListUserPluginResponseBodyResultBingoPlugins) GoString() string {
	return s.String()
}

func (s *ListUserPluginResponseBodyResultBingoPlugins) GetDescription() *string {
	return s.Description
}

func (s *ListUserPluginResponseBodyResultBingoPlugins) GetElasticsearchVersion() *string {
	return s.ElasticsearchVersion
}

func (s *ListUserPluginResponseBodyResultBingoPlugins) GetFileVersion() *string {
	return s.FileVersion
}

func (s *ListUserPluginResponseBodyResultBingoPlugins) GetName() *string {
	return s.Name
}

func (s *ListUserPluginResponseBodyResultBingoPlugins) GetSource() *string {
	return s.Source
}

func (s *ListUserPluginResponseBodyResultBingoPlugins) GetState() *string {
	return s.State
}

func (s *ListUserPluginResponseBodyResultBingoPlugins) GetVersion() *string {
	return s.Version
}

func (s *ListUserPluginResponseBodyResultBingoPlugins) SetDescription(v string) *ListUserPluginResponseBodyResultBingoPlugins {
	s.Description = &v
	return s
}

func (s *ListUserPluginResponseBodyResultBingoPlugins) SetElasticsearchVersion(v string) *ListUserPluginResponseBodyResultBingoPlugins {
	s.ElasticsearchVersion = &v
	return s
}

func (s *ListUserPluginResponseBodyResultBingoPlugins) SetFileVersion(v string) *ListUserPluginResponseBodyResultBingoPlugins {
	s.FileVersion = &v
	return s
}

func (s *ListUserPluginResponseBodyResultBingoPlugins) SetName(v string) *ListUserPluginResponseBodyResultBingoPlugins {
	s.Name = &v
	return s
}

func (s *ListUserPluginResponseBodyResultBingoPlugins) SetSource(v string) *ListUserPluginResponseBodyResultBingoPlugins {
	s.Source = &v
	return s
}

func (s *ListUserPluginResponseBodyResultBingoPlugins) SetState(v string) *ListUserPluginResponseBodyResultBingoPlugins {
	s.State = &v
	return s
}

func (s *ListUserPluginResponseBodyResultBingoPlugins) SetVersion(v string) *ListUserPluginResponseBodyResultBingoPlugins {
	s.Version = &v
	return s
}

func (s *ListUserPluginResponseBodyResultBingoPlugins) Validate() error {
	return dara.Validate(s)
}
