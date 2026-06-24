// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPluginAnalysisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *PluginAnalysisRequest
	GetBody() *string
	SetDryRun(v string) *PluginAnalysisRequest
	GetDryRun() *string
}

type PluginAnalysisRequest struct {
	// Request body parameters
	//
	// ```json
	//
	// example:
	//
	// {
	//
	//     "name": "plugin_name.zip",// plugin name
	//
	//     "ossObject": {
	//
	//       "bucketName": "bucketName",// oss bucket name
	//
	//       "key": "my_plugin_dir/plugin_name.zip" // oss file name
	//
	//     }
	//
	//   }
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
	// Whether to perform pre-upgrade validation:
	//
	// - true: Validate only
	//
	// - false: Upload directly
	//
	// example:
	//
	// false
	DryRun *string `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s PluginAnalysisRequest) String() string {
	return dara.Prettify(s)
}

func (s PluginAnalysisRequest) GoString() string {
	return s.String()
}

func (s *PluginAnalysisRequest) GetBody() *string {
	return s.Body
}

func (s *PluginAnalysisRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *PluginAnalysisRequest) SetBody(v string) *PluginAnalysisRequest {
	s.Body = &v
	return s
}

func (s *PluginAnalysisRequest) SetDryRun(v string) *PluginAnalysisRequest {
	s.DryRun = &v
	return s
}

func (s *PluginAnalysisRequest) Validate() error {
	return dara.Validate(s)
}
