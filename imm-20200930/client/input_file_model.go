// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInputFile interface {
	dara.Model
	String() string
	GoString() string
	SetAddresses(v []*Address) *InputFile
	GetAddresses() []*Address
	SetAlbum(v string) *InputFile
	GetAlbum() *string
	SetAlbumArtist(v string) *InputFile
	GetAlbumArtist() *string
	SetArtist(v string) *InputFile
	GetArtist() *string
	SetComposer(v string) *InputFile
	GetComposer() *string
	SetContentType(v string) *InputFile
	GetContentType() *string
	SetCustomId(v string) *InputFile
	GetCustomId() *string
	SetCustomLabels(v map[string]interface{}) *InputFile
	GetCustomLabels() map[string]interface{}
	SetFigures(v []*InputFileFigures) *InputFile
	GetFigures() []*InputFileFigures
	SetFileHash(v string) *InputFile
	GetFileHash() *string
	SetLabels(v []*Label) *InputFile
	GetLabels() []*Label
	SetLatLong(v string) *InputFile
	GetLatLong() *string
	SetMediaType(v string) *InputFile
	GetMediaType() *string
	SetOSSURI(v string) *InputFile
	GetOSSURI() *string
	SetPerformer(v string) *InputFile
	GetPerformer() *string
	SetProduceTime(v string) *InputFile
	GetProduceTime() *string
	SetTitle(v string) *InputFile
	GetTitle() *string
	SetURI(v string) *InputFile
	GetURI() *string
}

type InputFile struct {
	Addresses    []*Address             `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	Album        *string                `json:"Album,omitempty" xml:"Album,omitempty"`
	AlbumArtist  *string                `json:"AlbumArtist,omitempty" xml:"AlbumArtist,omitempty"`
	Artist       *string                `json:"Artist,omitempty" xml:"Artist,omitempty"`
	Composer     *string                `json:"Composer,omitempty" xml:"Composer,omitempty"`
	ContentType  *string                `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	CustomId     *string                `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	Figures      []*InputFileFigures    `json:"Figures,omitempty" xml:"Figures,omitempty" type:"Repeated"`
	FileHash     *string                `json:"FileHash,omitempty" xml:"FileHash,omitempty"`
	Labels       []*Label               `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	LatLong      *string                `json:"LatLong,omitempty" xml:"LatLong,omitempty"`
	MediaType    *string                `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	OSSURI       *string                `json:"OSSURI,omitempty" xml:"OSSURI,omitempty"`
	Performer    *string                `json:"Performer,omitempty" xml:"Performer,omitempty"`
	ProduceTime  *string                `json:"ProduceTime,omitempty" xml:"ProduceTime,omitempty"`
	Title        *string                `json:"Title,omitempty" xml:"Title,omitempty"`
	URI          *string                `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s InputFile) String() string {
	return dara.Prettify(s)
}

func (s InputFile) GoString() string {
	return s.String()
}

func (s *InputFile) GetAddresses() []*Address {
	return s.Addresses
}

func (s *InputFile) GetAlbum() *string {
	return s.Album
}

func (s *InputFile) GetAlbumArtist() *string {
	return s.AlbumArtist
}

func (s *InputFile) GetArtist() *string {
	return s.Artist
}

func (s *InputFile) GetComposer() *string {
	return s.Composer
}

func (s *InputFile) GetContentType() *string {
	return s.ContentType
}

func (s *InputFile) GetCustomId() *string {
	return s.CustomId
}

func (s *InputFile) GetCustomLabels() map[string]interface{} {
	return s.CustomLabels
}

func (s *InputFile) GetFigures() []*InputFileFigures {
	return s.Figures
}

func (s *InputFile) GetFileHash() *string {
	return s.FileHash
}

func (s *InputFile) GetLabels() []*Label {
	return s.Labels
}

func (s *InputFile) GetLatLong() *string {
	return s.LatLong
}

func (s *InputFile) GetMediaType() *string {
	return s.MediaType
}

func (s *InputFile) GetOSSURI() *string {
	return s.OSSURI
}

func (s *InputFile) GetPerformer() *string {
	return s.Performer
}

func (s *InputFile) GetProduceTime() *string {
	return s.ProduceTime
}

func (s *InputFile) GetTitle() *string {
	return s.Title
}

func (s *InputFile) GetURI() *string {
	return s.URI
}

func (s *InputFile) SetAddresses(v []*Address) *InputFile {
	s.Addresses = v
	return s
}

func (s *InputFile) SetAlbum(v string) *InputFile {
	s.Album = &v
	return s
}

func (s *InputFile) SetAlbumArtist(v string) *InputFile {
	s.AlbumArtist = &v
	return s
}

func (s *InputFile) SetArtist(v string) *InputFile {
	s.Artist = &v
	return s
}

func (s *InputFile) SetComposer(v string) *InputFile {
	s.Composer = &v
	return s
}

func (s *InputFile) SetContentType(v string) *InputFile {
	s.ContentType = &v
	return s
}

func (s *InputFile) SetCustomId(v string) *InputFile {
	s.CustomId = &v
	return s
}

func (s *InputFile) SetCustomLabels(v map[string]interface{}) *InputFile {
	s.CustomLabels = v
	return s
}

func (s *InputFile) SetFigures(v []*InputFileFigures) *InputFile {
	s.Figures = v
	return s
}

func (s *InputFile) SetFileHash(v string) *InputFile {
	s.FileHash = &v
	return s
}

func (s *InputFile) SetLabels(v []*Label) *InputFile {
	s.Labels = v
	return s
}

func (s *InputFile) SetLatLong(v string) *InputFile {
	s.LatLong = &v
	return s
}

func (s *InputFile) SetMediaType(v string) *InputFile {
	s.MediaType = &v
	return s
}

func (s *InputFile) SetOSSURI(v string) *InputFile {
	s.OSSURI = &v
	return s
}

func (s *InputFile) SetPerformer(v string) *InputFile {
	s.Performer = &v
	return s
}

func (s *InputFile) SetProduceTime(v string) *InputFile {
	s.ProduceTime = &v
	return s
}

func (s *InputFile) SetTitle(v string) *InputFile {
	s.Title = &v
	return s
}

func (s *InputFile) SetURI(v string) *InputFile {
	s.URI = &v
	return s
}

func (s *InputFile) Validate() error {
	if s.Addresses != nil {
		for _, item := range s.Addresses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Figures != nil {
		for _, item := range s.Figures {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InputFileFigures struct {
	FigureClusterId *string `json:"FigureClusterId,omitempty" xml:"FigureClusterId,omitempty"`
	FigureId        *string `json:"FigureId,omitempty" xml:"FigureId,omitempty"`
	FigureType      *string `json:"FigureType,omitempty" xml:"FigureType,omitempty"`
}

func (s InputFileFigures) String() string {
	return dara.Prettify(s)
}

func (s InputFileFigures) GoString() string {
	return s.String()
}

func (s *InputFileFigures) GetFigureClusterId() *string {
	return s.FigureClusterId
}

func (s *InputFileFigures) GetFigureId() *string {
	return s.FigureId
}

func (s *InputFileFigures) GetFigureType() *string {
	return s.FigureType
}

func (s *InputFileFigures) SetFigureClusterId(v string) *InputFileFigures {
	s.FigureClusterId = &v
	return s
}

func (s *InputFileFigures) SetFigureId(v string) *InputFileFigures {
	s.FigureId = &v
	return s
}

func (s *InputFileFigures) SetFigureType(v string) *InputFileFigures {
	s.FigureType = &v
	return s
}

func (s *InputFileFigures) Validate() error {
	return dara.Validate(s)
}
