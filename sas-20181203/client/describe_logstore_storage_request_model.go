// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogstoreStorageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *DescribeLogstoreStorageRequest
	GetFrom() *string
	SetLang(v string) *DescribeLogstoreStorageRequest
	GetLang() *string
}

type DescribeLogstoreStorageRequest struct {
	// The ID of the request source. Set the value to **sas**.
	//
	// This parameter is required.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
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
}

func (s DescribeLogstoreStorageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogstoreStorageRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogstoreStorageRequest) GetFrom() *string {
	return s.From
}

func (s *DescribeLogstoreStorageRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeLogstoreStorageRequest) SetFrom(v string) *DescribeLogstoreStorageRequest {
	s.From = &v
	return s
}

func (s *DescribeLogstoreStorageRequest) SetLang(v string) *DescribeLogstoreStorageRequest {
	s.Lang = &v
	return s
}

func (s *DescribeLogstoreStorageRequest) Validate() error {
	return dara.Validate(s)
}
