// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuildCacheConfig interface {
	dara.Model
	String() string
	GoString() string
	SetKeyPath(v map[string]interface{}) *BuildCacheConfig
	GetKeyPath() map[string]interface{}
	SetPaths(v []*string) *BuildCacheConfig
	GetPaths() []*string
}

type BuildCacheConfig struct {
	// example:
	//
	// { 	"3C75C832-0EAD-40D6-8FA1-2BA9171C926B": "~/.npm", 	"D256BB7A-1886-4A19-A75B-A1FDC23D5A00": "~/.cache" }
	KeyPath map[string]interface{} `json:"keyPath,omitempty" xml:"keyPath,omitempty"`
	Paths   []*string              `json:"paths,omitempty" xml:"paths,omitempty" type:"Repeated"`
}

func (s BuildCacheConfig) String() string {
	return dara.Prettify(s)
}

func (s BuildCacheConfig) GoString() string {
	return s.String()
}

func (s *BuildCacheConfig) GetKeyPath() map[string]interface{} {
	return s.KeyPath
}

func (s *BuildCacheConfig) GetPaths() []*string {
	return s.Paths
}

func (s *BuildCacheConfig) SetKeyPath(v map[string]interface{}) *BuildCacheConfig {
	s.KeyPath = v
	return s
}

func (s *BuildCacheConfig) SetPaths(v []*string) *BuildCacheConfig {
	s.Paths = v
	return s
}

func (s *BuildCacheConfig) Validate() error {
	return dara.Validate(s)
}
