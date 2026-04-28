// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScanFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*File) *ScanFileResponseBody
	GetItems() []*File
	SetNextMarker(v string) *ScanFileResponseBody
	GetNextMarker() *string
}

type ScanFileResponseBody struct {
	// The information about the files.
	Items []*File `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s ScanFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ScanFileResponseBody) GoString() string {
	return s.String()
}

func (s *ScanFileResponseBody) GetItems() []*File {
	return s.Items
}

func (s *ScanFileResponseBody) GetNextMarker() *string {
	return s.NextMarker
}

func (s *ScanFileResponseBody) SetItems(v []*File) *ScanFileResponseBody {
	s.Items = v
	return s
}

func (s *ScanFileResponseBody) SetNextMarker(v string) *ScanFileResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ScanFileResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
