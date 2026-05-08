// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextThemeListResult interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TextThemeListResult
	GetRequestId() *string
	SetTextThemeList(v []*TextTheme) *TextThemeListResult
	GetTextThemeList() []*TextTheme
}

type TextThemeListResult struct {
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	TextThemeList []*TextTheme `json:"textThemeList,omitempty" xml:"textThemeList,omitempty" type:"Repeated"`
}

func (s TextThemeListResult) String() string {
	return dara.Prettify(s)
}

func (s TextThemeListResult) GoString() string {
	return s.String()
}

func (s *TextThemeListResult) GetRequestId() *string {
	return s.RequestId
}

func (s *TextThemeListResult) GetTextThemeList() []*TextTheme {
	return s.TextThemeList
}

func (s *TextThemeListResult) SetRequestId(v string) *TextThemeListResult {
	s.RequestId = &v
	return s
}

func (s *TextThemeListResult) SetTextThemeList(v []*TextTheme) *TextThemeListResult {
	s.TextThemeList = v
	return s
}

func (s *TextThemeListResult) Validate() error {
	if s.TextThemeList != nil {
		for _, item := range s.TextThemeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
