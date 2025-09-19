// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteStrategyRequest
	GetId() *string
	SetLang(v string) *DeleteStrategyRequest
	GetLang() *string
	SetSourceIp(v string) *DeleteStrategyRequest
	GetSourceIp() *string
}

type DeleteStrategyRequest struct {
	// The ID of the baseline check policy that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1404656
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 192.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStrategyRequest) GoString() string {
	return s.String()
}

func (s *DeleteStrategyRequest) GetId() *string {
	return s.Id
}

func (s *DeleteStrategyRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteStrategyRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DeleteStrategyRequest) SetId(v string) *DeleteStrategyRequest {
	s.Id = &v
	return s
}

func (s *DeleteStrategyRequest) SetLang(v string) *DeleteStrategyRequest {
	s.Lang = &v
	return s
}

func (s *DeleteStrategyRequest) SetSourceIp(v string) *DeleteStrategyRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteStrategyRequest) Validate() error {
	return dara.Validate(s)
}
