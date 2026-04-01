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
	// example:
	//
	// {\\"totalCount\\": 1, \\"X-Total-Count\\": 1}
	Headers map[string]interface{} `json:"Headers,omitempty" xml:"Headers,omitempty"`
	// example:
	//
	// 5EEF8FAE-EEDD***
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListUserPluginResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
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
	BingoPlugins []*ListUserPluginResponseBodyResultBingoPlugins `json:"bingoPlugins,omitempty" xml:"bingoPlugins,omitempty" type:"Repeated"`
	// example:
	//
	// ct-test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// USER
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// UNINSTALLED
	State *string `json:"state,omitempty" xml:"state,omitempty"`
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
	// example:
	//
	// The plugin***
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 8.17.0
	ElasticsearchVersion *string `json:"elasticsearchVersion,omitempty" xml:"elasticsearchVersion,omitempty"`
	// example:
	//
	// CAEQaRiBgIDI2tie6hkiIGIwM2I3MjZmNjk3YzR***
	FileVersion *string `json:"fileVersion,omitempty" xml:"fileVersion,omitempty"`
	// example:
	//
	// dynamic-name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// USER
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// UNINSTALLED
	State *string `json:"state,omitempty" xml:"state,omitempty"`
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
