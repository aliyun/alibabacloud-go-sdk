// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTargetSubtitle interface {
	dara.Model
	String() string
	GoString() string
	SetDisableSubtitle(v bool) *TargetSubtitle
	GetDisableSubtitle() *bool
	SetExtractSubtitle(v *TargetSubtitleExtractSubtitle) *TargetSubtitle
	GetExtractSubtitle() *TargetSubtitleExtractSubtitle
	SetStream(v []*int32) *TargetSubtitle
	GetStream() []*int32
}

type TargetSubtitle struct {
	// Specifies whether to disable subtitle generation. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// >  If you call the GenerateVideoPlaylist operation and subtitles are required, you must set this parameter to false.
	//
	// example:
	//
	// false
	DisableSubtitle *bool `json:"DisableSubtitle,omitempty" xml:"DisableSubtitle,omitempty"`
	// The subtitle extraction settings.
	//
	// >  The GenerateVideoPlaylist operation does not support this parameter.
	ExtractSubtitle *TargetSubtitleExtractSubtitle `json:"ExtractSubtitle,omitempty" xml:"ExtractSubtitle,omitempty" type:"Struct"`
	// The index numbers of subtitle streams that need to be processed. If you set this parameter to null (default) or a value greater than 100, all subtitle streams are processed.
	//
	// 	- For example, you can set the parameter to `[0,1]` to process subtitle streams with index numbers 0 and 1, `[1]` to process only the subtitle stream with the index number 1, and `[101]` to process all subtitle streams.
	//
	// >  If you specify an index number but no subtitle stream with the index number is found, the index number is ignored.
	Stream []*int32 `json:"Stream,omitempty" xml:"Stream,omitempty" type:"Repeated"`
}

func (s TargetSubtitle) String() string {
	return dara.Prettify(s)
}

func (s TargetSubtitle) GoString() string {
	return s.String()
}

func (s *TargetSubtitle) GetDisableSubtitle() *bool {
	return s.DisableSubtitle
}

func (s *TargetSubtitle) GetExtractSubtitle() *TargetSubtitleExtractSubtitle {
	return s.ExtractSubtitle
}

func (s *TargetSubtitle) GetStream() []*int32 {
	return s.Stream
}

func (s *TargetSubtitle) SetDisableSubtitle(v bool) *TargetSubtitle {
	s.DisableSubtitle = &v
	return s
}

func (s *TargetSubtitle) SetExtractSubtitle(v *TargetSubtitleExtractSubtitle) *TargetSubtitle {
	s.ExtractSubtitle = v
	return s
}

func (s *TargetSubtitle) SetStream(v []*int32) *TargetSubtitle {
	s.Stream = v
	return s
}

func (s *TargetSubtitle) Validate() error {
	if s.ExtractSubtitle != nil {
		if err := s.ExtractSubtitle.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TargetSubtitleExtractSubtitle struct {
	// The format of the extracted subtitle file. Valid values:
	//
	// 	- ass
	//
	// 	- srt
	//
	// 	- webvtt
	//
	// example:
	//
	// webvtt
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The prefix of the OSS URI where the extracted subtitles are stored. The OSS URI is in the oss://bucket/object format, where bucket specifies the name of the OSS bucket that is in the same region as the current project and object specifies the full file path that includes the file name extension.
	//
	// 	- Example: If the prefix is oss://examplebucket/outputSubtitle, an output subtitle file has a URI in the format of oss://examplebucket/outputSubitile_${index}.${ext}. In the URI format, ${ext} is the file name extension of the output subtitle file, and ${index} is the same 0-based index number as that of the corresponding source subtitle stream file.
	//
	// example:
	//
	// oss://test-bucket/extractsubtitle
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s TargetSubtitleExtractSubtitle) String() string {
	return dara.Prettify(s)
}

func (s TargetSubtitleExtractSubtitle) GoString() string {
	return s.String()
}

func (s *TargetSubtitleExtractSubtitle) GetFormat() *string {
	return s.Format
}

func (s *TargetSubtitleExtractSubtitle) GetURI() *string {
	return s.URI
}

func (s *TargetSubtitleExtractSubtitle) SetFormat(v string) *TargetSubtitleExtractSubtitle {
	s.Format = &v
	return s
}

func (s *TargetSubtitleExtractSubtitle) SetURI(v string) *TargetSubtitleExtractSubtitle {
	s.URI = &v
	return s
}

func (s *TargetSubtitleExtractSubtitle) Validate() error {
	return dara.Validate(s)
}
