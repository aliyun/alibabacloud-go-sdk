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
	// The addresses.
	Addresses []*Address `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// The album.
	//
	// example:
	//
	// FirstAlbum
	Album *string `json:"Album,omitempty" xml:"Album,omitempty"`
	// The album artist.
	//
	// example:
	//
	// Jane
	AlbumArtist *string `json:"AlbumArtist,omitempty" xml:"AlbumArtist,omitempty"`
	// The artist.
	//
	// example:
	//
	// Jane
	Artist *string `json:"Artist,omitempty" xml:"Artist,omitempty"`
	// The composer.
	//
	// example:
	//
	// Jane
	Composer *string `json:"Composer,omitempty" xml:"Composer,omitempty"`
	// In most cases, you can leave this parameter empty. The Multipurpose Internet Mail Extensions (MIME) type of the file.
	//
	// example:
	//
	// image/jpeg
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The custom ID of the file. This parameter is optional. When the metadata of the file is indexed into the dataset, the custom ID is stored as the data attribute. You can map the custom ID to other data in your business system. You can configure this parameter based on your business requirements. For example, you can associate a URI with an ID in your business system. We recommend that you set this parameter to a unique value.
	//
	// This parameter supports prefix searches and sorting during queries. For more information, see [Supported fields and operators](https://help.aliyun.com/document_detail/252856.html).
	//
	// example:
	//
	// member-image-id-0001
	CustomId *string `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	// The custom labels of the file. This parameter is optional. The parameter stores custom key-value labels, which can be used to filter data. For more information, see [Supported fields and operators](https://help.aliyun.com/document_detail/252856.html).
	//
	// example:
	//
	// {
	//
	//       "MemberName": "Tim",
	//
	//       "Enabled": "True",
	//
	//       "ItemCount": "10"
	//
	// }
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	// This parameter is optional. The persons. This parameter is used to remove a face from a face group or modify a face group. For more information, see [Face clustering](https://help.aliyun.com/document_detail/477175.html).
	//
	// >  This parameter takes effect only for the UpdateFileMeta or BatchUpdateFileMeta operation.
	Figures []*InputFileFigures `json:"Figures,omitempty" xml:"Figures,omitempty" type:"Repeated"`
	// The file hash. In most cases, you can leave this parameter empty. This parameter is required only when the URI parameter specifies a file in Photo and Drive Service.
	//
	// example:
	//
	// 1d9c280a7c4f67f7ef873e28449dbe17
	FileHash *string `json:"FileHash,omitempty" xml:"FileHash,omitempty"`
	// The intelligent labels.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The GPS latitude and longitude information.
	//
	// example:
	//
	// 30.134390,120.074997
	LatLong *string `json:"LatLong,omitempty" xml:"LatLong,omitempty"`
	// In most cases, you can leave this parameter empty. The media type of the file.
	//
	// Enumerated values:
	//
	// 	- image
	//
	// 	- other
	//
	// 	- document
	//
	// 	- archive
	//
	// 	- video
	//
	// 	- audio
	//
	// example:
	//
	// image
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The path of the OSS object. In most cases, you can leave this parameter empty. You can specify this parameter only if the URI parameter specifies a file in Photo and Drive Service.
	//
	// example:
	//
	// oss://test-bucket/test-object.jpg
	OSSURI *string `json:"OSSURI,omitempty" xml:"OSSURI,omitempty"`
	// The performer.
	//
	// example:
	//
	// Jane
	Performer *string `json:"Performer,omitempty" xml:"Performer,omitempty"`
	// The time when the image was taken.
	//
	// example:
	//
	// 2021-06-29T14:50:13.011643661+08:00
	ProduceTime *string `json:"ProduceTime,omitempty" xml:"ProduceTime,omitempty"`
	// The file title.
	//
	// example:
	//
	// test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The URI of the file for which you want to create or update an index in the request. This parameter is required. The URI can represent an object in Object Storage Service (OSS) or a file in Photo and Drive Service.
	//
	// The OSS URI must be in the oss://${Bucket}/${Object} format. `${Bucket}` specifies the name of the OSS bucket that is in the same region as the current project. `${Object}` specifies the full file path that contains the object name extension.
	//
	// The URI of a file in Photo and Drive Service must be in the `pds://domains/${domain}/drives/${drive}/files/${file}/revisions/${revision}` format.
	//
	// >  URIs that start with HTTP are not supported.
	//
	// example:
	//
	// oss://examplebucket/sampleobject.jpg
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
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
	// The ID of the face cluster. The following IDs of special face clusters are reserved:
	//
	// 	- figure-cluster-id-independent: indicates that the face does not belong to any face cluster. The face may be added to a face cluster in subsequent face clustering tasks after new images are added to the dataset.
	//
	// 	- figure-cluster-id-unavailable: indicates that the face has not been included in a face clustering task since a new image was added to the dataset.
	//
	// example:
	//
	// Cluster-dbe72fec-b84c-4ab6-885b-3678e64****
	FigureClusterId *string `json:"FigureClusterId,omitempty" xml:"FigureClusterId,omitempty"`
	// The person ID.
	//
	// example:
	//
	// 2cb3c51e-b406-4b0c-af1b-897d88e1****
	FigureId *string `json:"FigureId,omitempty" xml:"FigureId,omitempty"`
	// The figure type. Set this parameter to `face`.
	//
	// example:
	//
	// face
	FigureType *string `json:"FigureType,omitempty" xml:"FigureType,omitempty"`
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
