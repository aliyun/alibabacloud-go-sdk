// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLibraryListResult interface {
	dara.Model
	String() string
	GoString() string
	SetLibrarys(v []*Library) *LibraryListResult
	GetLibrarys() []*Library
	SetRequestId(v string) *LibraryListResult
	GetRequestId() *string
}

type LibraryListResult struct {
	Librarys  []*Library `json:"librarys,omitempty" xml:"librarys,omitempty" type:"Repeated"`
	RequestId *string    `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s LibraryListResult) String() string {
	return dara.Prettify(s)
}

func (s LibraryListResult) GoString() string {
	return s.String()
}

func (s *LibraryListResult) GetLibrarys() []*Library {
	return s.Librarys
}

func (s *LibraryListResult) GetRequestId() *string {
	return s.RequestId
}

func (s *LibraryListResult) SetLibrarys(v []*Library) *LibraryListResult {
	s.Librarys = v
	return s
}

func (s *LibraryListResult) SetRequestId(v string) *LibraryListResult {
	s.RequestId = &v
	return s
}

func (s *LibraryListResult) Validate() error {
	if s.Librarys != nil {
		for _, item := range s.Librarys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
