// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DescribeEventDetailRequest
	GetId() *int64
	SetLang(v string) *DescribeEventDetailRequest
	GetLang() *string
}

type DescribeEventDetailRequest struct {
	// The ID of the anomalous event.
	//
	// > You can call the **DescribeEvents*	- operation to query the ID of the anomalous event.
	//
	// This parameter is required.
	//
	// example:
	//
	// 13456723343
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeEventDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventDetailRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeEventDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeEventDetailRequest) SetId(v int64) *DescribeEventDetailRequest {
	s.Id = &v
	return s
}

func (s *DescribeEventDetailRequest) SetLang(v string) *DescribeEventDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeEventDetailRequest) Validate() error {
	return dara.Validate(s)
}
