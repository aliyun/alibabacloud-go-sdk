// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddressResp interface {
	dara.Model
	String() string
	GoString() string
	SetImportAddress(v []*GetAddressResp) *ListAddressResp
	GetImportAddress() []*GetAddressResp
	SetNextMarker(v string) *ListAddressResp
	GetNextMarker() *string
	SetTruncated(v bool) *ListAddressResp
	GetTruncated() *bool
}

type ListAddressResp struct {
	ImportAddress []*GetAddressResp `json:"ImportAddress,omitempty" xml:"ImportAddress,omitempty" type:"Repeated"`
	// example:
	//
	// <your-next-address-name>
	NextMarker *string `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	// example:
	//
	// true
	Truncated *bool `json:"Truncated,omitempty" xml:"Truncated,omitempty"`
}

func (s ListAddressResp) String() string {
	return dara.Prettify(s)
}

func (s ListAddressResp) GoString() string {
	return s.String()
}

func (s *ListAddressResp) GetImportAddress() []*GetAddressResp {
	return s.ImportAddress
}

func (s *ListAddressResp) GetNextMarker() *string {
	return s.NextMarker
}

func (s *ListAddressResp) GetTruncated() *bool {
	return s.Truncated
}

func (s *ListAddressResp) SetImportAddress(v []*GetAddressResp) *ListAddressResp {
	s.ImportAddress = v
	return s
}

func (s *ListAddressResp) SetNextMarker(v string) *ListAddressResp {
	s.NextMarker = &v
	return s
}

func (s *ListAddressResp) SetTruncated(v bool) *ListAddressResp {
	s.Truncated = &v
	return s
}

func (s *ListAddressResp) Validate() error {
	if s.ImportAddress != nil {
		for _, item := range s.ImportAddress {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
